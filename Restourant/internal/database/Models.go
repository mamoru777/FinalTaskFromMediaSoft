package database

import (
	"github.com/google/uuid"
	"time"
)

type Menu struct {
	Id              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OnDate          time.Time `db:"on_date" gorm:"on_date"`
	OpeningRecordAt time.Time `db:"opening_record_at" gorm:"opening_record_at"`
	ClosingRecordAt time.Time `db:"closing_record_at" gorm:"closing_record_at"`
	CreatedAt       time.Time `db:"created_at" gorm:"created_at"`
	Products        []Product `gorm:"many2many:menu_product;"`
}
type Product struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string
	Description string
	Weight      int32
	Price       float64
	CreatedAt   time.Time `db:"created_at" gorm:"created_at"`
	Menus       []Menu    `gorm:"many2many:menu_product;"`
	Orders      []Order
	ProductType string
	//ProductType   ProductType `db:"product_type_id" gorm:"foreignKey:product_type_id"`
}

type ProductType struct {
	Id              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductTypeType string    `db:"product_type_type" gorm:"product_type_type"`
	//Products        []Product
}

type Order struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductName string    `db:"product_name" gorm:"product_name"`
	Count       int64     `db:"count" gorm:"count"`
	ProductId   uint
	Product     Product `db:"product_id" gorm:"foreignKey:product_id"`
}
