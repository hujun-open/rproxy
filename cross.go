package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

type CrossConnection struct {
	ID           int
	Conn1, Conn2 *net.TCPConn
}

func (cc CrossConnection) String() string {
	return fmt.Sprintf("crossconnection %d between %v and %v", cc.ID, cc.Conn1.RemoteAddr(), cc.Conn2.RemoteAddr())
}
func (cc *CrossConnection) Complete(c2 *net.TCPConn) error {
	if cc.Conn2 != nil {
		return fmt.Errorf("%v is already completed", cc)
	}
	cc.Conn2 = c2
	return nil
}

func (cc *CrossConnection) Run() {
	if cc.Conn1 == nil || cc.Conn2 == nil {
		log.Printf("%v is not fully initialized", cc.String())
		return
	}
	log.Printf("%v started", cc.String())
	defer log.Printf("%v ended", cc.String())
	wg2 := new(sync.WaitGroup)
	wg2.Add(2)
	go func() {
		n, err := io.Copy(cc.Conn1, cc.Conn2)
		log.Printf("%v-> %v ended, %d bytes copied, err is %v",
			cc.Conn2.RemoteAddr(), cc.Conn1.RemoteAddr(), n, err)
		cc.Conn1.Close()
		cc.Conn2.Close()
		wg2.Done()

	}()
	go func() {
		n, err := io.Copy(cc.Conn2, cc.Conn1)
		log.Printf("%v-> %v ended, %d bytes copied, err is %v",
			cc.Conn1.RemoteAddr(), cc.Conn2.RemoteAddr(), n, err)
		cc.Conn1.Close()
		cc.Conn2.Close()
		wg2.Done()
	}()
	wg2.Wait()
}
