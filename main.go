package main

import (
	"assignment_order/database"
	"assignment_order/router"
)

// @title       Orders API
// @version     1.0
// @description This is assignment 2 hacktiv8.

// @contact.name  Aditya Gocendra
// @contact.email gocendra123@gmail.com

// @host localhost:8080/orders/

func main() {
	database.StartDB()

	router.StartServer().Run(":8080")
}