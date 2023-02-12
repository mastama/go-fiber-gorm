package controller

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	db "github.com/mastama/go-fiber-gorm/config"
	"github.com/mastama/go-fiber-gorm/model"
)

func CreateCashier(c *fiber.Ctx) error {

	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Invalid data",
				"data":            data,
			})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Cashier name is required!",
				"data":            data,
			})
	}

	if data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Cashier passcode is required!",
				"data":            "Please insert your passcode",
			})
	}

	//saving cashier to db
	cashier := model.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	db.DB.Create(&cashier)
	return c.Status(200).JSON(
		fiber.Map{
			"responseCode":    true,
			"responseMessage": "Cashier added successfully",
			"data":            cashier,
		})
}

func EditCashier(c *fiber.Ctx) error {
	return nil
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func CashiersList(c *fiber.Ctx) error {
	var cashier []model.Cashier
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	return c.Status(200).JSON(
		fiber.Map{
			"responseCode":    true,
			"responseMessage": "Cashier list api",
			"data":            cashier,
		})
}

func GetCashierDetails(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
