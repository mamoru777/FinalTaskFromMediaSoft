package database

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}
func (s *Storage) CreateMenu(ctx context.Context, m *Menu) error {
	return s.db.WithContext(ctx).Create(m).Error
}

/*
	func (s *Storage) GetMenuId(ctx context.Context, created_at timestamp.Timestamp) (uuid.UUID, error) {
		m := new(Menu)
		err := s.db.WithContext(ctx).First(m, created_at).Error
		return m.Id, err
	}
*/
func (s *Storage) GetMenu(ctx context.Context, onDate time.Time) (*Menu, error) {
	m := new(Menu)
	err := s.db.Preload("Products").Where("EXTRACT(year FROM on_date) = ? AND EXTRACT(month FROM on_date) = ? AND EXTRACT(day FROM on_date) = ?", onDate.Year(), onDate.Month(), onDate.Day()).First(&m).Error
	return m, err
}

func (s *Storage) UpdateMenu(ctx context.Context, m *Menu) error {
	return s.db.WithContext(ctx).Save(m).Error
}

func (s *Storage) DeleteMenu(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Delete(&Menu{Id: id}).Error
}

func (s *Storage) CreateProduct(ctx context.Context, p *Product) error {
	return s.db.WithContext(ctx).Create(p).Error
}

func (s *Storage) GetProduct(ctx context.Context, name string) Product {
	p := *new(Product)
	s.db.WithContext(ctx).Where("name = ?", name).First(&p)
	return p
}

func (s *Storage) GetProductList(ctx context.Context) ([]*Product, error) {
	pl := []*Product{}
	err := s.db.WithContext(ctx).Find(&pl).Error
	return pl, err
}

func (s *Storage) UpdateProduct(ctx context.Context, p *Product) error {
	return s.db.WithContext(ctx).Save(p).Error
}

func (s *Storage) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Delete(&Product{Id: id}).Error
}

/*func (s *Storage) GetProductTypeList(ctx context.Context) ([]*ProductType, error) {
	ptl := []*ProductType{}
	err := s.db.WithContext(ctx).Find(ptl, &ProductType{}).Error
	return ptl, err
}*/
