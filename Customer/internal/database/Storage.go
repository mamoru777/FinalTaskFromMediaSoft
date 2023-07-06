package database

import (
	"database/sql"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateCustomer(ctx context.Context, c *Customer) error {
	const q = `
			insert into customers (name, office_id, office_name, created_at)
				values ($1, $2, $3, $4)
	`
	err := s.db.QueryRowContext(ctx, q, c.Name, c.OfficeId, c.OfficeName, c.CreatedAt).Err()
	return err
}

func (s *Storage) GetCustomerList(ctx context.Context, officeid uuid.UUID) ([]*Customer, error) {
	const q = `
			select id, name, office_id, office_name, created_at from customers where office_id = $1
	`
	var list []*Customer
	rows, err := s.db.QueryContext(ctx, q, officeid)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := new(Customer)
		err := rows.Scan(&c.Id, &c.Name, &c.OfficeId, &c.OfficeName, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}

func (s *Storage) CreateOffice(ctx context.Context, o *Office) error {
	const q = `
			insert into offices (name, address, created_at)
				values ($1, $2, $3)
	`
	err := s.db.QueryRowContext(ctx, q, o.Name, o.Address, o.CreatedAt).Err()
	return err
}

func (s *Storage) GetOfficeList(ctx context.Context) ([]*Office, error) {
	const q = `
			SELECT * FROM offices
	`
	var list []*Office
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		o := new(Office)
		err := rows.Scan(&o.Id, &o.Name, &o.Address, &o.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, o)

	}
	return list, nil
}
