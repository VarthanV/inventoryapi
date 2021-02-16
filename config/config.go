package config

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

type DBConfig struct{
	Host string
	Port int
	User string
	DBName string
	Password string
}

func initEnv(){
	if err := godotenv.Load(); err != nil {
        panic("No .env file found")
    }
}

func BuildDBConfig()*DBConfig{
	initEnv()
	dbConfig := DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: 3306,
		DBName: os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USERNAME"),	
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
	 "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
	 dbConfig.User,
	 dbConfig.Password,
	 dbConfig.Host,
	 dbConfig.Port,
	 dbConfig.DBName,
	)
   }