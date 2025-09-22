package main

import (
	microfiber "github.com/MegaBytee/micro-fiber"
	"github.com/MegaBytee/micro-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := microfiber.Config{
		AuthKeyLookup: "header:apiKey",
		Port:          "3690",
		Cache:         false,
		Limitter:      true,
		Logger:        true,
		Metrics:       true,
	}
	service := microfiber.NewService(&config)

	hello := routes.NewRoute(routes.GET, "/", func(c *fiber.Ctx) error {
		return c.JSON(routes.NewResponseHTTP(true, "hello", "hello world"))
	})

	protected := routes.NewRoute(routes.GET, "/protected", func(c *fiber.Ctx) error {
		return c.JSON(routes.NewResponseHTTP(true, "protected", "hello from protected"))
	}).SetProtected(true)

	//log.Println("protected:", protected.Protected)
	routes := []*routes.ApiRoute{hello, protected}

	service.RegisterRoutes(routes)
	service.Setup()
	service.Start()
}
