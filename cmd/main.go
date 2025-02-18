package main

import (
	"github.com/Gaurav-coding08/persistence-go/cmd/server"
	"github.com/Gaurav-coding08/persistence-go/config"
	"github.com/Gaurav-coding08/persistence-go/database/connect"
)

func main() {
	cfg := config.LoadConfig()

	connect.InitDB(cfg)
	db := connect.DB

	server.StartConsumers(cfg, db)

	// Block forever to prevent exit
	select {}
}
