package main

import (
	"FinalTaskFromMediaSoft/Customer/internal"
	"FinalTaskFromMediaSoft/Customer/internal/configdb"
	"FinalTaskFromMediaSoft/Customer/internal/configserv"
	"github.com/caarlos0/env/v8"
	"log"
)

func main() {
	cfg := configdb.ConfigDb{}
	config := configserv.ConfigServ{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to retrieve env variables, %v", err)
	}
	if err := env.Parse(&config); err != nil {
		log.Fatal("failed to retrieve env variables, %v", err)
	}
	if err := internal.Run(cfg, config); err != nil {
		log.Fatal("error running server ", err)
	}
}
