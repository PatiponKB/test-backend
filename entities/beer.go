package entities

import (
	"time"
	_beerModel "github.com/PatiponKB/backend-test/pkg/beer/model"
)
type Beer struct {
	ID          	uint64 			`gorm:"primaryKey;autoIncrement"`
	Name        	string			`gorm:"type:varchar(64);unique;not null;"`
	Type        	string			`gorm:"type:varchar(128);"`
	Description 	string			`gorm:"type:varchar(128);"`
	Picture     	string			`gorm:"type:varchar(256);not null;"`
	CreatedAt   	time.Time		`gorm:"not null;autoCreateTime;"`
	UpdatedAt		time.Time		`gorm:"not null;autoUpdateTime;"`
}

func (b *Beer) ToModel() *_beerModel.Beer {
	return &_beerModel.Beer{
		ID: b.ID,
		Name: b.Name,
		Type: b.Type,
		Description: b.Description,
		Picture: b.Picture,
	}
}