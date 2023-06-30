package database

import (
	"github.com/google/uuid"
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
func (s *Storage) CreateMenu(ctx context.Context, m *Menu) error {
	return s.db.WithContext(ctx).Create(m).Error
}

func (s *Storage) GetMenu(ctx context.Context, id uuid.UUID) (*Menu, error) {
	m := new(Menu)
	err := s.db.WithContext(ctx).First(m, id).Error
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

func (s *Storage) GetProduct(ctx context.Context, id uuid.UUID) (*Product, error) {
	p := new(Product)
	err := s.db.WithContext(ctx).First(p, id).Error
	return p, err
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
