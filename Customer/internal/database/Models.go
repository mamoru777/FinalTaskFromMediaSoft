package database

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	//Id       uint   `db:"id" gorm:"id;primaryKey;type:serial;unique_index;auto_increment;sequence:customer_id_seq"`
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name       string    `db:"name" gorm:"name"`
	CreatedAt  time.Time
	OfficeId   uuid.UUID
	OfficeName string
	Office     Office `db:"office_id" gorm:"foreignKey:office_id"`
	Orders     []Order
}

type Office struct {
	//Id        uint      `db:"id" gorm:"id;primaryKey;type:serial;unique_index;auto_increment;sequence:office_id_seq"`
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `db:"name" gorm:"name"`
	Address   string    `db:"address" gorm:"address"`
	CreatedAt time.Time `db:"created_at" gorm:"created_at"`
	Customers []Customer
}

type Order struct {
	//Id         uint `db:"id" gorm:"id;primaryKey;type:serial;unique_index;auto_increment;sequence:order_id_seq"`
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CustomerId uint
	Customer   Customer    `db:"customer_id" gorm:"foreignKey:customer_id"`
	OrderItems []OrderItem `gorm:"many2many:order_orderitems;"`
}

type OrderItem struct {
	//Id     uint    `db:"id" gorm:"id;primaryKey;type:serial;unique_index;auto_increment;sequence:orderitem_id_seq"`
	Id     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Count  int64     `db:"count" gorm:"count"`
	Orders []Order   `gorm:"many2many:order_orderitems;"`
}

type Menu struct {
	Id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Products []Product `gorm:"many2many:menu_product;"`
}

type Product struct {
	Id                    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                  string    `db:"name" gorm:"name"`
	Description           string    `db:"description" gorm:"description"`
	Weight                int64     `db:"weight" gorm:"weight"`
	Price                 int64     `db:"price" gorm:"price"`
	CreatedAt             time.Time `db:"created_at" gorm:"created_at"`
	Menus                 []Menu    `gorm:"many2many:menu_product;"`
	CustomerProductTypeId uint
	CustomerProductType   CustomerProductType `db:"customer_product_type_id" gorm:"foreignKey:customer_product_type_id"`
}
type CustomerProductType struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductType string    `db:"product_type" gorm:"product_type"`
	Products    []Product
}
