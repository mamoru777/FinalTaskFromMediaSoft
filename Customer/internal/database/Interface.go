package database

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Rep interface {
	CreateCustomer(ctx context.Context, c *Customer) error
	GetCustomerList(ctx context.Context, officeid uuid.UUID) ([]*Customer, error)

	CreateOffice(ctx context.Context, o *Office) error
	GetOfficeList(ctx context.Context) ([]*Office, error)
}
