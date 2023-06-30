package database

import (
	"FinalTaskFromMediaSoft/Customer/internal/configdb"
	"FinalTaskFromMediaSoft/Customer/internal/model"
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
	db.AutoMigrate(&model.Office{}, &model.Customer{}, &model.Order{}, &model.OrderItem{}, &model.CustomerProductType{}, &model.Menu{}, &model.Product{})
	return db, err
	/*db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err = db.Exec("CREATE DATABASE dbcustomer").Error
		if err != nil {
			log.Fatal("Failed to Create DataBase")
		} else {
			fmt.Print("Success to Create to DataBase")
		}
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		fmt.Print("Success to connect to DataBase")
	}
	db.AutoMigrate(&model.Customer{}, &model.Office{})
	return db, err*/
}
