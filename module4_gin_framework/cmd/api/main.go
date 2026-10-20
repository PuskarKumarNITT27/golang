package main

import (
	"fmt"
	"gin-crud/internal/config"
	"gin-crud/internal/db"
	"gin-crud/internal/server"
	"log"
)

func main() {
	cfg, err := config.Load();
	if err != nil {
		log.Fatal("Config Error")
	}

	client,database,err := db.Connect(cfg)

	if err != nil {
		log.Fatal("Database Connection Error")
	}

	defer func(){
		if err := db.Disconnect(client) ; err != nil {
			log.Printf("Mongo Disconnect error")
		}

	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s",cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatal("Server Failed")
	}
}