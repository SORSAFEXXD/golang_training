package controllers

import (
	"errors"
	"strings"
	"time"

	"last_project/database"
	"last_project/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetProfiles(c *fiber.Ctx) error {
	var profiles []models.UserProfile
	if err := database.DBConn.Order("id").Find(&profiles).Error; err != nil {
		return databaseError(c, err)
	}
	return c.JSON(profiles)
}

func GetProfile(c *fiber.Ctx) error {
	employeeID, err := uuid.Parse(c.Params("employee_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid employee_id"})
	}

	var profile models.UserProfile
	if err := database.DBConn.First(&profile, "employee_id = ?", employeeID).Error; err != nil {
		return recordError(c, err)
	}
	return c.JSON(profile)
}

func SearchProfiles(c *fiber.Ctx) error {
	search := strings.TrimSpace(c.Query("search"))
	if search == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "search is required"})
	}

	var profiles []models.UserProfile
	pattern := "%" + strings.ToLower(search) + "%"
	result := database.DBConn.Order("id").Where(
		"employee_id = ? OR LOWER(name) LIKE ? OR LOWER(lastname) LIKE ?",
		search, pattern, pattern,
	).Find(&profiles)
	if result.Error != nil {
		return databaseError(c, result.Error)
	}
	return c.JSON(profiles)
}

func CreateProfile(c *fiber.Ctx) error {
	var input models.ProfileInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validateProfileInput(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if input.EmployeeID == uuid.Nil {
		input.EmployeeID = uuid.New()
	}
	age, err := calculateProfileAge(input.Birthday, time.Now())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid birthday"})
	}
	profile := models.UserProfile{
		EmployeeID: input.EmployeeID,
		Name:       strings.TrimSpace(input.Name),
		Lastname:   strings.TrimSpace(input.Lastname),
		Birthday:   input.Birthday,
		Age:        age,
		Email:      strings.TrimSpace(input.Email),
		Tel:        strings.TrimSpace(input.Tel),
	}
	if err := database.DBConn.Create(&profile).Error; err != nil {
		return databaseError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	employeeID, err := uuid.Parse(c.Params("employee_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid employee_id"})
	}

	var profile models.UserProfile
	if err := database.DBConn.First(&profile, "employee_id = ?", employeeID).Error; err != nil {
		return recordError(c, err)
	}

	var input models.ProfileInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validateProfileInput(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	if input.EmployeeID != uuid.Nil && input.EmployeeID != employeeID {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "employee_id cannot be changed"})
	}
	age, err := calculateProfileAge(input.Birthday, time.Now())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid birthday"})
	}

	profile.Name = strings.TrimSpace(input.Name)
	profile.Lastname = strings.TrimSpace(input.Lastname)
	profile.Birthday = input.Birthday
	profile.Age = age
	profile.Email = strings.TrimSpace(input.Email)
	profile.Tel = strings.TrimSpace(input.Tel)
	if err := database.DBConn.Save(&profile).Error; err != nil {
		return databaseError(c, err)
	}
	return c.JSON(profile)
}

func DeleteProfile(c *fiber.Ctx) error {
	employeeID, err := uuid.Parse(c.Params("employee_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid employee_id"})
	}
	result := database.DBConn.Where("employee_id = ?", employeeID).Delete(&models.UserProfile{})
	if result.Error != nil {
		return databaseError(c, result.Error)
	}
	if result.RowsAffected == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.SendStatus(fiber.StatusOK)
}

func validateProfileInput(input models.ProfileInput) error {
	if strings.TrimSpace(input.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(input.Lastname) == "" {
		return errors.New("lastname is required")
	}
	if strings.TrimSpace(input.Birthday) == "" {
		return errors.New("birthday is required")
	}
	birthday, err := time.Parse("2006-01-02", input.Birthday)
	if err != nil {
		return errors.New("birthday must use YYYY-MM-DD format")
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if birthday.After(today) {
		return errors.New("birthday cannot be in the future")
	}
	return nil
}

func GetAgeGroups(c *fiber.Ctx) error {
	var profiles []models.UserProfile
	if err := database.DBConn.Find(&profiles).Error; err != nil {
		return databaseError(c, err)
	}

	counts := map[string]int{
		"GenZ": 0, "GenY": 0, "GenX": 0, "Baby Boomer": 0, "G.I. Generation": 0,
	}
	for _, profile := range profiles {
		counts[ageGroupFor(profile.Age)]++
	}
	return c.JSON(fiber.Map{
		"count":     len(profiles),
		"age_count": counts,
	})
}

func ageGroupFor(age int) string {
	switch {
	case age < 24:
		return "GenZ"
	case age <= 41:
		return "GenY"
	case age <= 56:
		return "GenX"
	case age <= 75:
		return "Baby Boomer"
	default:
		return "G.I. Generation"
	}
}

func calculateProfileAge(birthday string, now time.Time) (int, error) {
	birth, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, err
	}

	age := now.Year() - birth.Year()
	if now.Month() < birth.Month() || (now.Month() == birth.Month() && now.Day() < birth.Day()) {
		age--
	}
	return age, nil
}

func recordError(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return databaseError(c, err)
}

func databaseError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
}
