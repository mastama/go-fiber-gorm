package route

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/mastama/go-fiber-gorm/controller"
)

func Setup(app *fiber.App) {

	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logout)
	app.Post("/cashiers/:cashierId/passcode", controller.Passcode)

	//cashier routes
	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/cashiers", controller.CashiersList)
	app.Get("/cashiers/:cashierId", controller.GetCashierDetails)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)

}
