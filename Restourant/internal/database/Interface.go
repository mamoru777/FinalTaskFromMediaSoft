package database

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Rep interface {
	CreateMenu(ctx context.Context, m *Menu) error
	GetMenu(ctx context.Context, onDate time.Time) (*Menu, error)
	UpdateMenu(ctx context.Context, m *Menu) error
	DeleteMenu(ctx context.Context, id uuid.UUID) error

	CreateProduct(ctx context.Context, p *Product) error
	GetProduct(ctx context.Context, name string) Product
	GetProductList(ctx context.Context) ([]*Product, error)
	UpdateProduct(ctx context.Context, p *Product) error
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	GetProductById(ctx context.Context, id uuid.UUID) string
	//CreateProductType(ctx context.Context, pt *ProductType) error
	//GetProductType(ctx context.Context, id int64) (*ProductType, error)
	//GetProductTypeList(ctx context.Context) ([]*ProductType, error)
	//UpdateProductType(ctx context.Context, pt *ProductType) error
	//DeleteProductType(ctx context.Context, id int64) error

	CreateOrder(ctx context.Context, o *Order) error
	GetOrdersByCustomer(ctx context.Context, customerId uuid.UUID) ([]*Order, error)
	//UpdateOrder(ctx context.Context, o *Order) error
	//DeleteOrder(ctx context.Context, id int64) error*/
	GetOrderList(ctx context.Context) ([]*Order, error)
}
