package main

import (
	"fmt"
	"log"
	"flag"
	"net"
	"google.golang.org/grpc"
	pb "warikan.suwageeks.org/warikan/pkg/pb"
)

// commandLine variable
var (
	port = flag.Int("port", 50000, "The gRPC Server Port")
)

// wrap
type server struct {
	pb.UnimplementedWarikanServer
}

// main
func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to Listen:\n\t %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWarikanServer(s, &server{})
	log.Printf("Start Server Listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve:\n\t%v", err)
	}
}
