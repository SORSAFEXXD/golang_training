package main

import (
	"fmt"
	"log"

	"day3/database"
	"day3/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	m "day3/models"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"1234",
		"127.0.0.1",
		"3307",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	if err := database.DBConn.AutoMigrate(&m.Dogs{}, &m.Company{}); err != nil {
		panic(err)
	}
}

func main() {
	app := fiber.New()
	routes.InetRoutes(app)
	initDatabase()

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
