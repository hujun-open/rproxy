package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"rproxy/api"
	"runtime"
	"time"

	"github.com/elazarl/goproxy"
)

const (
	defaultToClientListenPort = 7777
	defaultToWOrkerListenPort = 7778
	defaultReflAPIListenPort  = 7779
	defaultProxyPort          = 8080
	workerRole                = "worker"
	reflRole                  = "refl"
)

func main() {
	lcport := flag.Uint("clport", defaultToClientListenPort, "http client facing listen port")
	lwport := flag.Uint("wlport", defaultToWOrkerListenPort, "worker facing listen port")
	apiport := flag.Uint("apiport", defaultReflAPIListenPort, "reflector API listen port")
	role := flag.String("role", workerRole, "role")
	refladdr := flag.String("refl", "", "reflector tcp address")
	reflapiaddr := flag.String("reflapi", "", "reflector api tcp address")
	svraddr := flag.String("svr", "", "server tcp address")
	proxyPort := flag.Uint("proxyport", defaultProxyPort, "http proxy listen port")
	localProxy := flag.Bool("localproxy", true, "use local http proxy")
	profiling := flag.Bool("p", false, "enable profiling")
	flag.Parse()
	if *profiling {
		runtime.SetBlockProfileRate(1000000000)
		go func() {
			log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()

	}
	switch *role {
	default:
		log.Fatalf("invalid role %v", *role)
	case reflRole:
		refl, err := NewReflector(fmt.Sprintf("0.0.0.0:%d", *lcport),
			fmt.Sprintf("0.0.0.0:%d", *lwport), int(*apiport))
		if err != nil {
			log.Fatal(err)
		}
		go refl.ListenForWorker()
		refl.ListenForClient()
	case workerRole:
		if *localProxy {
			proxy := goproxy.NewProxyHttpServer()
			proxy.Verbose = true
			*svraddr = fmt.Sprintf("127.0.0.1:%d", *proxyPort)
			go func() {
				log.Fatal(http.ListenAndServe(*svraddr, proxy))
			}()
			log.Print("starting local proxy server")
			time.Sleep(3 * time.Second)
		}
		worker, err := NewWorker(context.Background(), *reflapiaddr, *refladdr, *svraddr)
		if err != nil {
			log.Fatal(err)
		}
		defer worker.clnt.Signoff(context.Background(), &api.Empty{})
		go worker.reportToRefl()
		worker.listenForCreateReq()
	}
}
