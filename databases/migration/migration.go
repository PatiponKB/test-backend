package main

import (
	"context"
	"fmt"

	"github.com/PatiponKB/backend-test/config"
	"github.com/PatiponKB/backend-test/databases"
	"github.com/PatiponKB/backend-test/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewMariaDatabase(conf.MariaDB)

	tx := db.Connection().Begin()

	mongoClient := databases.ConnectMongoDB()
	defer mongoClient.Disconnect(context.Background())
	fmt.Println(mongoClient)

	beerMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
	
}


func beerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Beer{})
}