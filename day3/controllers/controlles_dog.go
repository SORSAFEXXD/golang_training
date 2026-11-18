package controllers

import (
	"day3/database"
	m "day3/models"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// DOG CRUD
func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDeletedDogs(c *fiber.Ctx) error {
	var dogs []m.Dogs
	result := database.DBConn.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dogs)
}

func GetDogsByIDRange(c *fiber.Ctx) error {
	var dogs []m.Dogs
	result := database.DBConn.Where("dog_id > ? AND dog_id < ?", 50, 100).Find(&dogs)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	if result := db.Find(&dogs); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error.Error(),
		})
	}

	var dataResults []m.DogsRes
	sumColor := map[string]int{
		"red":      0,
		"green":    0,
		"pink":     0,
		"no color": 0,
	}
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}
		sumColor[typeStr]++

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	r := m.ResultData{
		Name:     "golang-test(dog)",
		Data:     dataResults,
		Sum:      len(dogs),
		SumColor: sumColor,
	}
	return c.Status(200).JSON(r)
}
