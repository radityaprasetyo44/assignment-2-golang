package main

import (
	"assignment2/configs"
	"assignment2/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs.LoadEnv()

	routes.LoadRoute()
}
