package cmd

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wiyco/go-api-template/pkg/api"
	H "github.com/wiyco/go-api-template/pkg/handler"
)

func InitApi() {
	app := fiber.New(fiber.Config{
		ErrorHandler: H.HandleAPIError,
	})

	api.InitRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
