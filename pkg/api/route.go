package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiyco/go-api-template/pkg/api/v1/ping"
)

func InitRoutes(app *fiber.App) {
	apiv1 := app.Group("/api/v1")

	ping.InitPingRoutes(apiv1)
}
