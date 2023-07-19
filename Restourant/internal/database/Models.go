package database

import (
	"github.com/google/uuid"
	"time"
)

type Menu struct {
	Id              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OnDate          time.Time `db:"on_date" gorm:"on_date;type:timestamp without time zone"`
	OpeningRecordAt time.Time `db:"opening_record_at" gorm:"opening_record_at;type:timestamp without time zone"`
	ClosingRecordAt time.Time `db:"closing_record_at" gorm:"closing_record_at;type:timestamp without time zone"`
	CreatedAt       time.Time `db:"created_at" gorm:"created_at;type:timestamp without time zone"`
	Products        []Product `gorm:"many2many:menu_product;"`
}
type Product struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
	Weight      int32
	Price       float64
	CreatedAt   time.Time `db:"created_at" gorm:"created_at;type:timestamp without time zone"`
	Menus       []Menu    `gorm:"many2many:menu_product;"`
	Orders      []Order
	ProductType string
}

type Order struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductName string    `db:"product_name" gorm:"product_name"`
	Count       int64     `db:"count" gorm:"count"`
	ProductId   uuid.UUID
	CustomerId  uuid.UUID `gorm:"customer_id"`
	CreatedAt   time.Time `db:"created_at" gorm:"created_at;type:timestamp without time zone"`
	Product     Product   `db:"product_id" gorm:"foreignKey:product_id"`
}
