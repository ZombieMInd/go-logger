package main

import (
	"fmt"
	"log"

	"github.com/ZombieMInd/go-logger/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := &server.Config{}

	server.InitConfig(conf)

	fmt.Printf("Starting %s on %s \n", conf.Name, conf.BindAddr)
	server.Start(conf)
}
