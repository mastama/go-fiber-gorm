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

func UpdateCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier model.Cashier

	db.DB.Find(&cashier, "id=?", cashierId)

	//validation for checking cashierId
	if cashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Cashier not found",
			})
	}

	var UpdateCashier model.Cashier
	err := c.BodyParser(&UpdateCashier)
	if err != nil {
		return err
	}

	if UpdateCashier.Name == "" {
		return c.Status(404).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Cashier name is required",
			})
	}

	cashier.Name = UpdateCashier.Name
	cashier.Passcode = UpdateCashier.Passcode
	db.DB.Save(&cashier)

	return c.Status(404).JSON(
		fiber.Map{
			"responseCode":    true,
			"responseMessage": "Success",
			"data":            cashier,
		})
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
	cashierId := c.Params("cashierId")
	var cashier model.Cashier

	db.DB.Select("id, name, created_at, updated_at").Where("id =?", cashierId).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	//what if there is no cashier or cashier id not provided
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "Cashier not found",
				"error":           map[string]interface{}{},
			})
	}

	return c.Status(200).JSON(fiber.Map{
		"responseCode":    true,
		"responseMessage": "success",
		"data":            cashierData,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier model.Cashier

	db.DB.Where("id = ?", cashierId).First(&cashier)
	// its similar in mysql or sql query like
	// select * from cashier where id = cashierId and limit 1 etc
	if cashier.Id == 0 {
		return c.Status(404).JSON(
			fiber.Map{
				"responseCode":    false,
				"responseMessage": "cashier not found",
			})
	}

	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(404).JSON(
		fiber.Map{
			"responseCode":    true,
			"responseMessage": "Cashier deleted successfully",
		})
}
