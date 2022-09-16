package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type env struct {
	Port int

	MySQL struct {
		Port     int
		Host     string
		User     string
		Password string
		DBName   string
	}
}

var Env env

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Env not found")
	}

	convPort, _ := strconv.Atoi(os.Getenv("PORT"))
	Env.Port = convPort

	convMySqlPort, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	Env.MySQL.Port = convMySqlPort

	Env.MySQL.Host = os.Getenv("MYSQL_HOST")
	Env.MySQL.User = os.Getenv("MYSQL_USER")
	Env.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	Env.MySQL.DBName = os.Getenv("MYSQL_DB_NAME")
}
