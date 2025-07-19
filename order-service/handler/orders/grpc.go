package handler 

import (
	"context"
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	ordersService types.OrderService 
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService){
	grpcHandler := &OrderGrpcHandler{
		ordersService: ordersService,
	}
	orders.RegisterOrderServiceServer(grpc, grpcHandler)
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderId: 22,
		CustomerId: 1,
		ProductId: 4,
		Quantity: 5,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}


	res := &orders.CreateOrderResponse{
		Status: "success",
		Order: order,
	}
	return res, nil
}