package databases

import "gorm.io/gorm"

type Database interface {
	Connection() *gorm.DB
}





