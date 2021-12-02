package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"rproxy/api"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Reflector struct {
	api.UnimplementedRProxyAPIServer
	CrossConnections         map[int]*CrossConnection //key is CC ID
	toClient, toWorker       *net.TCPListener
	createWorkerCrossReqChan chan *api.CreateWorkerCrossReq
	WorkerAddr               net.IP
	workerLock               *sync.RWMutex
	currentCCID              int
	fromWorkerConnQ          map[int]*net.TCPConn //key is the remote tcp port
	fromWorkerConnQLock      *sync.RWMutex
}

const (
	createWorkerCrossReqChanDepth = 128
)

func (refl *Reflector) Signon(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	p, _ := peer.FromContext(ctx)
	refl.ResetWorker(p.Addr.(*net.TCPAddr).IP)
	return &api.Empty{}, nil
}
func (refl *Reflector) Signoff(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	refl.ResetWorker(nil)
	return &api.Empty{}, nil
}
func (refl *Reflector) CreateWorkerCross(_ *api.Empty, stream api.RProxyAPI_CreateWorkerCrossServer) error {
	defer log.Print("createworker routine ended")
	for req := range refl.createWorkerCrossReqChan {
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send create worker channel, %v", err)
			return err
		}
	}
	return nil
}

func (refl *Reflector) ReportWorkerCross(stream api.RProxyAPI_ReportWorkerCrossServer) error {
	defer log.Print("report worker routine ended")
	for {
		worker, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to recv from report worker channel, %v", err)
			return err
		}

		refl.fromWorkerConnQLock.Lock()
		if conn, ok := refl.fromWorkerConnQ[int(worker.Port)]; ok {
			refl.workerLock.RLock()
			if cc, ok := refl.CrossConnections[int(worker.ID)]; ok {
				err = cc.Complete(conn)
				if err != nil {
					log.Printf("%v is already completed, can't add conn %v",
						cc.String(), conn.RemoteAddr())
				} else {
					//start CC
					go cc.Run()
					//remove from waiting Q
					delete(refl.fromWorkerConnQ, int(worker.Port))
				}
			}
			refl.workerLock.RUnlock()
		}
		refl.fromWorkerConnQLock.Unlock()

	}
}
func (refl *Reflector) ResetWorker(w net.IP) {
	refl.workerLock.Lock()
	defer refl.workerLock.Unlock()
	refl.CrossConnections = make(map[int]*CrossConnection)
	refl.WorkerAddr = w
}
func (refl *Reflector) ListenForWorker() {
	for {
		newworkerc, err := refl.toWorker.AcceptTCP()
		if err != nil {
			log.Fatalf("failed to accept client conn, %v", err)
		}
		refl.fromWorkerConnQLock.Lock()
		refl.fromWorkerConnQ[newworkerc.RemoteAddr().(*net.TCPAddr).Port] = newworkerc
		refl.fromWorkerConnQLock.Unlock()
		log.Printf("got a new worker data connection %v", newworkerc.RemoteAddr())
	}
}

func (refl *Reflector) ListenForClient() {
	for {
		refl.workerLock.RLock()
		if refl.WorkerAddr == nil {
			refl.workerLock.RUnlock()
			continue
		}
		refl.workerLock.RUnlock()
		newclinetc, err := refl.toClient.AcceptTCP()
		if err != nil {
			log.Fatalf("failed to accept client conn, %v", err)
		}
		log.Printf("accepted new client connection from %v", newclinetc.RemoteAddr())
		newcc := &CrossConnection{
			ID:    refl.currentCCID,
			Conn1: newclinetc,
		}
		refl.currentCCID++
		refl.workerLock.Lock()
		refl.CrossConnections[newcc.ID] = newcc
		refl.workerLock.Unlock()
		workreq := &api.CreateWorkerCrossReq{
			ID: uint32(newcc.ID),
		}
		refl.createWorkerCrossReqChan <- workreq
	}
}

func NewReflector(clientListenAddr, workerListenAddr string, apiport int) (*Reflector, error) {
	caddr, err := net.ResolveTCPAddr("tcp", clientListenAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid client listen address %v, %w", clientListenAddr, err)
	}
	waddr, err := net.ResolveTCPAddr("tcp", workerListenAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid worker listen address %v, %w", workerListenAddr, err)
	}
	r := new(Reflector)
	r.toClient, err = net.ListenTCP("tcp", caddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create client listener %v, %w", clientListenAddr, err)
	}
	r.toWorker, err = net.ListenTCP("tcp", waddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create worker listener %v, %w", workerListenAddr, err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", apiport))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on API port: %w", err)
	}
	s := grpc.NewServer()
	api.RegisterRProxyAPIServer(s, r)
	log.Printf("API listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r.workerLock = new(sync.RWMutex)
	r.fromWorkerConnQLock = new(sync.RWMutex)
	r.WorkerAddr = nil
	r.createWorkerCrossReqChan = make(chan *api.CreateWorkerCrossReq, createWorkerCrossReqChanDepth)
	r.fromWorkerConnQ = make(map[int]*net.TCPConn)

	return r, nil
}
