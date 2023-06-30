package main

import (
	"FinalTaskFromMediaSoft/Customer/internal"
	"FinalTaskFromMediaSoft/Customer/internal/configdb"
	"github.com/caarlos0/env/v8"
	"log"
)

func main() {
	cfg := configdb.ConfigDb{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("failed to retrieve env variables, %v", err)
	}
	if err := internal.Run(cfg); err != nil {
		log.Fatal("error running server ", err)
	}
}
