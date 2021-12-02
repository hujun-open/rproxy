package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"rproxy/api"
	"sync"

	"google.golang.org/grpc"
)

type Worker struct {
	clnt              api.RProxyAPIClient
	svrAddr, reflAddr string
	creatCCStream     api.RProxyAPI_CreateWorkerCrossClient
	reportCCStream    api.RProxyAPI_ReportWorkerCrossClient
	CrossConnections  map[int]*CrossConnection //conn1 is to refl
	CCLock            *sync.RWMutex
	reportChan        chan *api.ReportWorkerCrossReq
}

const (
	reportChanDepth = 128
)

func NewWorker(ctx context.Context, reflmgmtaddr, refldataaddr, svr string) (*Worker, error) {
	conn, err := grpc.Dial(reflmgmtaddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	r := new(Worker)
	r.clnt = api.NewRProxyAPIClient(conn)
	r.svrAddr = svr
	r.reflAddr = refldataaddr
	r.creatCCStream, err = r.clnt.CreateWorkerCross(ctx, &api.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to create createworker stream, %w", err)
	}
	r.reportCCStream, err = r.clnt.ReportWorkerCross(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create reportworker stream, %w", err)
	}
	_, err = r.clnt.Signon(ctx, &api.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to sign on, %w", err)
	}
	r.CCLock = new(sync.RWMutex)
	r.CrossConnections = make(map[int]*CrossConnection)
	r.reportChan = make(chan *api.ReportWorkerCrossReq, reportChanDepth)
	log.Printf("worker created, with refl api %v,refl data %v and sever %v",
		reflmgmtaddr, refldataaddr, svr)
	return r, nil
}
func (w *Worker) reportToRefl() {
	defer log.Print("report worker routine ended")
	var err error
	for {
		report := <-w.reportChan
		err = w.reportCCStream.Send(report)
		if err != nil {
			log.Printf("can't report to reflector, %v", err)
		}

	}
}

func (w *Worker) listenForCreateReq() {
	defer log.Print("listen for create worker req routine ended")
	for {
		req, err := w.creatCCStream.Recv()
		if err != nil {
			log.Fatalf("faild to recv from create worker stream, %v", err)
		}
		svrconn, err := net.Dial("tcp", w.svrAddr)
		if err != nil {
			if err != nil {
				log.Printf("can't connect to proxy server %v, %v", w.svrAddr, err)
				continue
			}
		}
		reflconn, err := net.Dial("tcp", w.reflAddr)
		if err != nil {
			if err != nil {
				log.Printf("can't connect to reflector %v, %v", w.reflAddr, err)
				continue
			}
		}
		cross := &CrossConnection{
			ID:    int(req.ID),
			Conn1: reflconn.(*net.TCPConn),
			Conn2: svrconn.(*net.TCPConn),
		}
		go cross.Run()
		w.CCLock.Lock()
		w.CrossConnections[int(req.ID)] = cross
		w.CCLock.Unlock()
		w.reportChan <- &api.ReportWorkerCrossReq{
			ID:   req.ID,
			Port: uint32(reflconn.LocalAddr().(*net.TCPAddr).Port),
		}
	}
}
