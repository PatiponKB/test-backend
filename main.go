package main

import (
	"github.com/PatiponKB/backend-test/config"
	"github.com/PatiponKB/backend-test/databases"
	"github.com/PatiponKB/backend-test/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewMariaDatabase(conf.MariaDB)
	server := server.NewechoServer(conf, db.Connection())

	server.Start()
}