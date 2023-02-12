package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/mastama/go-fiber-gorm/model"
)

var DB *gorm.DB

func Connect() {
	//DB Connection
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connection := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("connection db failed!")
	}

	DB = db
	fmt.Println("connection db successfully")

	AutoMigrate(db)
	// END DB CONNECTION
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&model.Cashier{},
		&model.Category{},
		&model.Discount{},
		&model.Order{},
		&model.Payment{},
		&model.PaymentType{},
		&model.Product{},
	)
}
