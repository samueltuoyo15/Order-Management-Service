package main

func main(){
	grpcServer := NewGrpcServer(":8080")
	grpcServer.Run()
}