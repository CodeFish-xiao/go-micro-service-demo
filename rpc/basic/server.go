package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc/basic/string-service"
)

func main() {
	stringService := new(service.StringService)
	registerError := rpc.Register(stringService)
	if registerError != nil {
		log.Fatal("Register error: ", registerError)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
