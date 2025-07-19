package main

func main(){
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	grpcServer := NewGrpcServer(":9000")
	grpcServer.Run()
}