package route

import "github.com/gofiber/fiber/v2"

func RouteInit(r *fiber.App) {

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Hello": "Jancok!",
		})
	})

}
