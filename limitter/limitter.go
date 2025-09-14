package limitter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Setup(app *fiber.App) {

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))
}
