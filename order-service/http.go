package main 

import (
	"net/http"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/service"
	"github.com/samueltuoyo15/Order-Management-Service/utils"
	handler "github.com/samueltuoyo15/Order-Management-Service/order-service/handler/orders"

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
	orderHandler := handler.NewHttpOrdersHandler(orderService)
	orderHandler.RegisterRouter(router)
	logger.Info("Http Server started and listening on", "port", s.addr)
	
	return http.ListenAndServe(s.addr, router)
}