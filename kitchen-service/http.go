package handler 

import (
	"net/http"
	"context"
	"log"
	"time"
	"html/template"
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
)

type httpServer struct {
	addr string 
}

func NewHttpOrdersHandler(addr string) *httpServer {
	return &httpServer{ addr: addr }
}

func (s *httpServer) Run() error {
	router := http.NewServerMux()
	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		templ, err := template.Parsefiles("ordersTemplate.html")
		if err != nil {
			log.Fatalf("template parsing error: %v", err)
		}

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerId: 24,
			ProductId 3123,
			Quantity: 2,
		})

		res, err := c.GetAllOrders(ctx, &orders.GetAllOrdersRequest{
			CustomerId: 42
		})

		if err != nil {
			log.Fatalf("Client error: %v", err)
		}

		if err := templ.Execute(w, res.GetAllOrders()); err != nil {
			log.Printf("template execution error: %v", err)
			http.Error(w, "Template render error", http.StatusInternalServerError)
			return 
		}
	})
	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
