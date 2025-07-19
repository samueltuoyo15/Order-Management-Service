package main 

import (
	"net/http"
	"context"
	"log"
	"time"
	"html/template"
	"github.com/samueltuoyo15/Order-Management-Service/utils"
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
)

type httpServer struct {
	addr string 
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{ addr: addr }
}

type OrderView struct {
	OrderId int32
	CustomerId int32
	Quantity int32
	CreatedAt string
	UpdatedAt string
}

func (s *httpServer) Run() error {
	logger := utils.InitLogger(true)
	router := http.NewServeMux()
	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		templ, err := template.ParseFiles("kitchen-service/ordersTemplate.html")
		if err != nil {
			log.Fatalf("template parsing error: %v", err)
		}

		_, err = c.CreateOrder(ctx, &orders.CreateOrderRequest{
			TraceId: "trace-id-123",
			CustomerId: 24,
			ProductId: 3123,
			Quantity: 2,
		})

		if err != nil {
			log.Fatalf("Client error: %v", err)
		}
		
		res, err := c.GetAllOrders(ctx, &orders.GetAllOrdersRequest{
			CustomerId: 42,
		})

		if err != nil {
			log.Fatalf("Client error: %v", err)
		}

		var ordersForView []OrderView
		for _, o := range res.Orders {
			ordersForView = append(ordersForView, OrderView{
				OrderId: o.OrderId,
				CustomerId: o.CustomerId,
				Quantity: o.Quantity,
				CreatedAt: o.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
				UpdatedAt: o.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
			})
		}
		if err := templ.Execute(w, ordersForView); err != nil {
			log.Printf("template execution error: %v", err)
			http.Error(w, "Template render error", http.StatusInternalServerError)
			return 
		}
	})
	logger.Info("Starting server on", "port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
