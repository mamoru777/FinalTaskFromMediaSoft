package internal

import (
	"FinalTaskFromMediaSoft/Customer/internal/configdb"
	"FinalTaskFromMediaSoft/Customer/internal/database"
	"log"
)

func Run(cfg configdb.ConfigDb) error {
	_, err := database.InitDb(cfg)
	if err != nil {
		log.Fatal("Cannot to init the database", err)
	}
	return nil
}
