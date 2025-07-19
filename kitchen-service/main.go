pacakge main

import (
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCClient(addr string) *grpc.ClientConn{
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != il {
		log.Fatalf("failed to connect: %v", err)
	}
	return conn
}

func main(){
	http.Server := NewHttpServer(":3000")
	httpServer.Run()
}