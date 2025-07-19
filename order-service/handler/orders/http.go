package handler 

import (
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
	"github.com/samueltuoyo15/Order-Management-Service/common/util"
	"github.com/samueltuoyo15/Order-Management-Service/order-service/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.orderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler {
		ordersService: orderService,
	}

	return handler
}
func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	routerHandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest 
	err := util.ParseJSON(r, &req)
	if eerr != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return 
	}

	order := &orders.Order{
		order_id: 42,
		customer_id: req.GetCustomerID()
		product_id: req.GetProductID(),
		Quantity: req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return 
	}

	res ::= &orders.CreateOrderResponse{ status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}