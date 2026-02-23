package proxy

import "github.com/gofiber/fiber/v2"

type Proxy interface {
	Accept(key string) bool
	Proxy(c *fiber.Ctx) error
}
