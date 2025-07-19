package handler 

import (
	"time"
	"crypto/rand"
	"encoding/binary"
	mathrand "math/rand"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (h *OrderGrpcHandler) GetAllOrders(ctx context.Context, req *orders.GetAllOrdersRequest) (*orders.GetAllOrdersResponse, error) {
		o := h.ordersService.GetAllOrders(ctx)
		res := &orders.GetAllOrdersResponse{
			Orders: o,
		}
		return res, nil
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	now := timestamppb.New(time.Now())
	customerID := generateCryptoRandomID(1000) 
	productID := generateMathRandomID(10, 99)

	order := &orders.Order{
		OrderId: generateMathRandomID(10000, 99999),
		CustomerId: customerID,
		ProductId: productID,
		Quantity: 5,
		CreatedAt: now,
		UpdatedAt: now,
	}
	customerID++
	productID++
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

func generateCryptoRandomID(max int32) int32 {
	var b [4]byte
	_, err := rand.Read(b[:])
	if err != nil {
		return 1 
	}
	val := int32(binary.BigEndian.Uint32(b[:]))
	if val < 0 {
		val = -val
	}
	return val % max
}

func generateMathRandomID(min, max int32) int32 {
	mathrand.Seed(time.Now().UnixNano())
	return min + int32(mathrand.Intn(int(max-min+1)))
}
