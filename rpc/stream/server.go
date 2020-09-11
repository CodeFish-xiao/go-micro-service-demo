package main

import (
	"flag"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"net"
	pb "rpc/stream-pb"
	"rpc/stream/string-service"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	stringService := new(string_service.StringService)
	pb.RegisterStringServiceServer(grpcServer, stringService)
	grpcServer.Serve(lis)
}
