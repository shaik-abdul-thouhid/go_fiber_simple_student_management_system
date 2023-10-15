package main

import (
	"fmt"
	"log"
	"time"

	"go_algo/db"
	"go_algo/db/models"
	"go_algo/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func main() {

	docs.SwaggerInfo.Title = "Go Algorithm Package"
	docs.SwaggerInfo.Description = "This is a simple Student Management System"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "student.go"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
	})

	db.Connect()

	app.Get("/students/:id?", func(c *fiber.Ctx) error {
		dataMap := make(map[string]any)

		students := models.FetchAllStudents()

		if value := c.Params("id"); value != "" {
			fmt.Println("id", value)
			for _, j := range students {
				if j.UiD == value {
					dataMap["time"] = time.Now()
					dataMap["data"] = j
					return c.JSON(dataMap)
				}
			}

			return c.JSON(struct {
				status        bool
				error_message string
			}{status: false, error_message: "id not found"})
		}

		dataMap["time"] = time.Now()
		dataMap["data"] = students

		return c.JSON(dataMap)
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	if err := app.Listen(":80"); err != nil {
		log.Fatal("unable to start server", err)
	}
}
