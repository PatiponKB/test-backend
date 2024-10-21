package main

import (
	"github.com/PatiponKB/backend-test/config"
	"github.com/PatiponKB/backend-test/databases"
	"github.com/PatiponKB/backend-test/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewMariaDatabase(conf.MariaDB)

	tx := db.Connection().Begin()

	beerAdding(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}

func beerAdding(tx *gorm.DB) {
	beers := []entities.Beer{		
		{
			Name:        "Stout beer",
			Description: "It is more popular than Porter because it tastes better.cl",
			Type: 		 "10",
			Picture:     "https://pin.it/3Ve94lyUq",
		},
		{
			Name:        "Porter beer",
			Description: "Black beer, bitter and sweet, nowadays very popular.",
			Type: 		 "10",
			Picture:     "https://pin.it/5AneJt74Q",
		},
		{
			Name:        "Pilsner beer",
			Description: "The word Pilsner is known as a beer that has a specific taste.",
			Type: 		 "20",
			Picture:     "https://pin.it/50XHspIf8",
		},
		{
			Name:        "Light beer",
			Description: "is different from Pilsner style because it has lower alcohol content.",
			Type: 		 "20",
			Picture:     "https://pin.it/7bDW7lGVh",
		},
		{
			Name:        "Bock beer",
			Description: "This type of beer is dark in color, has a strong aroma and taste.",
			Type: 		 "20",
			Picture:     "https://pin.it/6Z1S0jEFp",
		},
	}

	tx.CreateInBatches(beers, len(beers))
}