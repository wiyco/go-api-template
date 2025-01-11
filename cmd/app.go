package cmd

import (
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/wiyco/go-api-template/pkg/api"
	H "github.com/wiyco/go-api-template/pkg/handler"
)

func InitApi() {
	app := fiber.New(fiber.Config{
		ErrorHandler: H.HandleAPIError,
	})

	api.InitRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // default port
	}

	slog.Info("starting server", "port", port)

	err := app.Listen(":8000")
	if err != nil {
		slog.Error("failed to start server",
			"port", port,
			"error", err,
		)
		os.Exit(1)
	}
}
