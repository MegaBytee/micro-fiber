package routes

import "github.com/gofiber/fiber/v2"

type ApiRoute struct {
	Method    Method
	Path      string
	Handler   func(c *fiber.Ctx) error
	Protected bool
}

func NewRoute(method Method, path string, handler func(c *fiber.Ctx) error) *ApiRoute {
	return &ApiRoute{
		Method:    method,
		Path:      path,
		Handler:   handler,
		Protected: false,
	}
}

func (r *ApiRoute) SetProtected(value bool) *ApiRoute {
	r.Protected = value
	return r
}
func (r *ApiRoute) Set(app *fiber.App) {
	switch r.Method {
	case GET:
		app.Get(r.Path, r.Handler)
	case POST:
		app.Post(r.Path, r.Handler)
	case DELETE:
		app.Delete(r.Path, r.Handler)
	case PUT:
		app.Put(r.Path, r.Handler)
	}
}
