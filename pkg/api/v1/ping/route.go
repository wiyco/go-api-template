package ping

import (
	"github.com/gofiber/fiber/v2"
)

func InitPingRoutes(router fiber.Router) {
	router.Get("/ping", GetPing)
}
