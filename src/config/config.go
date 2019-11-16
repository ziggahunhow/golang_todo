package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload" // load env
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	host := os.Getenv("DBHost")
	password := os.Getenv("DBPassword")
	user := os.Getenv("DBUser")
	port, _ := strconv.Atoi(os.Getenv("DBPort"))

	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   "todos",
		Password: password,
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
