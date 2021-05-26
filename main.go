package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "tag-service/proto"
	"tag-service/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "port")
	flag.Parse()
}
func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err:%v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Serve err: %v", err)
	}
}
