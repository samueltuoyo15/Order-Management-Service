package main 

import (
	"log"
	"net/http"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/service"
	"github.com/samueltuoyo15/Order-Management-Service/utils"
)
type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr }
}
func (s *httpServer) Run() error {
	logger := utils.InitLogger(true)
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	logger.Info("Http Server started and listening on", s.addr)
	
	return http.ListnAndServe(s.addr, router)
}