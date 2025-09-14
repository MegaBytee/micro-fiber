package cache

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func Setup(app *fiber.App) {
	app.Use(cache.New())
}
