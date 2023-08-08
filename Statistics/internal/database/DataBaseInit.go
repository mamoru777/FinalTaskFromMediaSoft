package database

import (
	"FinalTaskFromMediaSoft/Statistics/internal/configdb"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDb(cfg configdb.ConfigDb) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PgHost, cfg.PgUser, cfg.PgPwd, cfg.PgDBName, cfg.PgPort)
	//return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot to Connect DataBase", err)
	}
	db.AutoMigrate(&Order{})
	return db, err
}
