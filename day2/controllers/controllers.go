package controllers

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserData struct {
	Email        string `json:"email" validate:"required,email"`
	UserName     string `json:"uname" validate:"required,username"`
	Pass         string `json:"pass" validate:"required,password"`
	LineId       string `json:"lineid" validate:"required"`
	Phone        string `json:"phone" validate:"required,numeric,len=10"`
	BusinessType string `json:"businesstype" validate:"required,business_type"`
	WebUrl       string `json:"weburl" validate:"required,endswith=sogoodweb.com"`
}

func Calculate(ctx *fiber.Ctx) error {
	nStr := ctx.Params("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 0 {
		return ctx.Status(fiber.StatusBadRequest).SendString("Invalid number")
	}

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return ctx.SendString(fmt.Sprintf("%d!=%d", n, fact))
}

func ChangeWordToInt(ctx *fiber.Ctx) error {
	taxID := ctx.Query("tax_id")
	if taxID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Please provide tax_id query param")
	}

	var asciiList []string
	for _, ch := range taxID {
		asciiList = append(asciiList, strconv.Itoa(int(ch)))
	}

	result := ""
	for i, val := range asciiList {
		if i > 0 {
			result += " "
		}
		result += val
	}

	return ctx.SendString(result)
}

var userValidator = func() *validator.Validate {
	v := validator.New()
	if err := v.RegisterValidation("username", func(field validator.FieldLevel) bool {
		return IsValidUsername(field.Field().String())
	}); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("password", func(field validator.FieldLevel) bool {
		return IsValidPassword(field.Field().String())
	}); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("business_type", func(field validator.FieldLevel) bool {
		switch field.Field().String() {
		case "individual", "company":
			return true
		default:
			return false
		}
	}); err != nil {
		panic(err)
	}
	return v
}()

func IsValidUsername(uname string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, uname)
	return matched
}

func IsValidPassword(pass string) bool {
	return len(pass) >= 6 && len(pass) <= 20
}

func Register(ctx *fiber.Ctx) error {
	var user UserData

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "รูปแบบ JSON ไม่ถูกต้อง",
		})
	}

	err := userValidator.Struct(user)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if !errors.As(err, &validationErrors) {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  false,
				"message": "ไม่สามารถตรวจสอบข้อมูลได้",
			})
		}

		messages := map[string]string{
			"Email":        "กรุณากรอกอีเมลให้ถูกต้อง",
			"UserName":     "ชื่อผู้ใช้ต้องเป็นตัวอักษรภาษาอังกฤษ, ตัวเลข, _ หรือ - เท่านั้น",
			"Pass":         "รหัสผ่านต้องมีความยาว 6-20 ตัวอักษร",
			"LineId":       "กรุณากรอก Line ID",
			"Phone":        "เบอร์โทรศัพท์ต้องเป็นตัวเลข 10 หลัก",
			"BusinessType": "BusinessType ต้องเป็น individual หรือ company",
			"WebUrl":       "WebUrl ต้องลงท้ายด้วย sogoodweb.com",
		}
		var validationMessages []string
		for _, fieldError := range validationErrors {
			message, ok := messages[fieldError.StructField()]
			if !ok {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  false,
					"message": "ไม่สามารถตรวจสอบข้อมูลได้",
				})
			}
			validationMessages = append(validationMessages, message)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"errors": validationMessages,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  true,
		"message": "สมัครสมาชิกสำเร็จ!",
		"data":    user,
	})
}
