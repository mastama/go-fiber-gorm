package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	db "github.com/mastama/go-fiber-gorm/config"
	"github.com/mastama/go-fiber-gorm/model"
	"os"
	"strconv"
	"time"
)

func Login(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Invalid post request",
			})
	}

	//now checking postcode field is empty or not
	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Passcode is required",
				"error":           map[string]interface{}{},
			})
	}
	var cashier model.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "passcode not match",
				//error is not neccasary but may be use in future
				"error": map[string]interface{}{},
			})
	}

	if cashier.Passcode != data["passcode"] {
		return c.Status(401).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "passcode not match",
				//error is not neccasary but may be use in future
				"error": map[string]interface{}{},
			})
	}

	//here we are create jwt token if record match in db
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.Id)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), //1day
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(401).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Token Expired or Invalid",
			})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = tokenString

	return c.Status(200).JSON(
		fiber.Map{
			"responseCode":    true,
			"responseMessage": "Success",
			"data":            cashierData,
		})
}

func Logout(c *fiber.Ctx) error {
	return nil
}

func Passcode(c *fiber.Ctx) error {
	return nil
}
