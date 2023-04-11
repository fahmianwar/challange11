package main

import (
	"challange11/database"
	"challange11/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
