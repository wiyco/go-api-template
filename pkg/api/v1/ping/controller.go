package ping

import (
	"github.com/gofiber/fiber/v2"

	H "github.com/wiyco/go-api-template/pkg/handler"
)

func GetPing(c *fiber.Ctx) error {
	return H.HandleAPISuccess(c, "pong")
}
