package database

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Count      int64     `db:"count" gorm:"count"`
	ProductId  uuid.UUID `gorm:"product_id"`
	CustomerId uuid.UUID `gorm:"customer_id"`
	CreatedAt  time.Time `db:"created_at" gorm:"created_at;type:timestamp without time zone"`
}
