package handler 

import (
	"net/http"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
	"github.com/samueltuoyo15/Order-Management-Service/common/util"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler {
		ordersService: orderService,
	}

	return handler
}
func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest 
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return 
	}
	now := timestamppb.New(time.Now())

	order := &orders.Order{
		OrderId: 42,
		CustomerId: req.GetCustomerId(),
		ProductId: req.GetProductId(),
		Quantity: req.GetQuantity(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	res := &orders.CreateOrderResponse{ Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}