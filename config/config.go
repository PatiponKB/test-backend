package config

import (
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		MariaDB *MariaDB 	`mapstructure:"mariadb"`
		Server 	*Server 	`mapstructure:"server"`		 
	}

	Server struct {
		Port         int           `mapstructure:"port"`
		AllowOrigins []string      `mapstructure:"allowOrigins"`
		Timeout      time.Duration `mapstructure:"timeout"`
		BodyLimit    string        `mapstructure:"bodyLimit" `    
	}

	MariaDB struct {
		User     string 	`mapstructure:"user" `
		Password string 	`mapstructure:"password" `
		Host     string 	`mapstructure:"host" `
		Port     int    	`mapstructure:"port" `
		DBName   string 	`mapstructure:"dbname" `
		Charset  string 	`mapstructure:"charset" `
		ParseTime bool  	`mapstructure:"parseTime" `
		Loc      string 	`mapstructure:"loc" `
	}

)

var (
	once 			sync.Once
	configInstance 	*Config
)

func ConfigGetting() *Config{
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

		if err := viper.ReadInConfig(); 
			err != nil{
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); 
			err != nil{
			panic(err)
		}

	})

	return configInstance
}
