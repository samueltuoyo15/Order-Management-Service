package main 

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/samueltuoyo15/Order-Management-Service/utils"
	handler "github.com/samueltuoyo15/Order-Management-Service/order-service/handler/orders"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/service"
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
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(server, orderService)
	
	logger.Info("Grpc Server started and listeneing", "port", s.addr)

	return server.Serve(listener)
}