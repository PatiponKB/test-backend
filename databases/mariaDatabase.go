package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/PatiponKB/backend-test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mariaDatabase struct {
	*gorm.DB
}

var (
	once                    sync.Once
	mariaDatabaseInstance 	*mariaDatabase
)

func NewMariaDatabase(conf *config.MariaDB) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DBName, conf.Charset, conf.ParseTime, conf.Loc)

		conndb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		log.Printf("Connected to Database %s", conf.DBName)

		mariaDatabaseInstance = &mariaDatabase{conndb}
	})

	return mariaDatabaseInstance
}

func (db *mariaDatabase) Connection() *gorm.DB {
	return mariaDatabaseInstance.DB
}