package database

import (
	"assignment_order/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// host     = "localhost"
	// user     = "postgres"
	// password = "aditgocendra"
	// dbPort     = "5432"
	// dbName   = "orders_by"
	// db       *gorm.DB
	// err      error

	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbPort   = os.Getenv("PGPORT")
	dbName   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)


func StartDB()  {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.Debug().AutoMigrate( models.Orders{}, models.Items{},)
}

func GetDB() *gorm.DB{
	return db
}