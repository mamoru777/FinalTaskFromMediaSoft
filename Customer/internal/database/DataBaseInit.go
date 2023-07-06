package database

import (
	"FinalTaskFromMediaSoft/Customer/internal/configdb"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
)

func InitDb(cfg configdb.ConfigDb) (*sql.DB, error) {
	db, err := sql.Open("pgx", formatConnect(cfg))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

	/*dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",

		cfg.PgHost, cfg.PgUser, cfg.PgPwd, cfg.PgDBName, cfg.PgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot to Connect DataBase", err)
	}
	db.AutoMigrate(&model.Office{}, &model.Customer{}, &model.Order{}, &model.OrderItem{}, &model.CustomerProductType{}, &model.Menu{}, &model.Product{})
	return db, err
	*/

}

func formatConnect(cfg configdb.ConfigDb) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPwd, cfg.PgHost, cfg.PgPort, cfg.PgDBName,
	)

}
