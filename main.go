package main

import (
	"github.com/gofiber/fiber/v2"
	db "github.com/mastama/go-fiber-gorm/config"
	route "github.com/mastama/go-fiber-gorm/route"
)

func main() {
	// call connection to db
	db.Connect()

	app := fiber.New()
	// app.Use(app)

	route.Setup(app)

	app.Listen(":8082")

}
