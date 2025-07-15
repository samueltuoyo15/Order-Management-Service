package types

import (
	"context"
	"github.com/samueltuoyo15/Order-Management-Service/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(ctx context.Context, order *orders.Order) error
}