package main 

import (
	"net"
	"google.golang.org/grpc"
	"github.com/samueltuoyo15/Order-Management-Service/utils"
)

type grpcServer struct {
	addr string
}

func NewGrpcServer(addr string) *grpcServer {
	return &grpcServer{addr: addr}
}

func (s *grpcServer) Run() error {  
	logger := utils.InitLogger(true)
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	server := grpc.NewServer()
	
	
	logger.Info("Grpc Server started and listeneing on port %v", s.addr)

	return server.Serve(listener)
}