package database

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateOrder(ctx context.Context, o *Order) error {
	return s.db.WithContext(ctx).Create(o).Error
}

func (s *Storage) GetOrderList(ctx context.Context) ([]*Order, error) {
	ol := []*Order{}
	err := s.db.WithContext(ctx).Find(&ol).Error
	return ol, err
}
