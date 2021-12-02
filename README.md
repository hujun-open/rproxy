# Overview
rproxy is a reflective TCP proxy for specific use case like following:

`client --- rproxy_reflector ==FW== rproxy_worker --- server`

client and server are 3rd party TCP applications, client needs to talk to the server, however there is no direction network connection between them, AND the "==FW==" network blocks any connection initiated from left side => right (e.g firewall), in other words, it only allows connection initiated from right side => left;

rproxy is designed to facilitate client talking to server in such case.

## How does it work?

1. run rproxy as reflector at computer1
2. run rproxy as worker at computer2
3. config client so it creates TCP connection conn-1 to reflector 
4. reflector instructs worker to create two TCP connections:
    * conn-A: between reflector and worker
    * conn-B: between worker and server
5. reflector cross-connects conn-1 and conn-A and worker cross-connect conn-A and conn-B, so at this point client could exchange TCP message with server.

`client --conn-1-- rproxy_reflector ==conn-A== rproxy_worker --conn-B-- server`

## CLI
Usage of rproxy.exe:
```
  -apiport uint
        reflector API listen port (default 7779)
  -clport uint
        http client facing listen port (default 7777)
  -localproxy
        use local http proxy (default true)
  -p    enable profiling
  -proxyport uint
        http proxy listen port (default 8080)
  -refl string
        reflector tcp address
  -reflapi string
        reflector api tcp address
  -role string
        role (default "worker")
  -svr string
        server tcp address
  -wlport uint
        worker facing listen port (default 7778)
```
* run as reflector,client connects to the `clport` 8001

`rproxy -role refl -apiport 8000 -clport 8001 -wlport 8002`

* run as worker connects to built-in HTTP proxy server

`rproxy -role worker -reflapi 10.10.10.1:8000 -refl 10.10.10.1:8002`

* run as worker connects to an external server @ 172.16.1.1:3000

`rproxy -role worker -reflapi 10.10.10.1:8000 -refl 10.10.10.1:8001 -svr 172.16.1.1:3000`
