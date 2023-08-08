package database

import "golang.org/x/net/context"

type Rep interface {
	CreateOrder(ctx context.Context, o *Order) error
	GetOrderList(ctx context.Context) ([]*Order, error)
}
