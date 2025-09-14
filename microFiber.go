package microfiber

import (
	"log"

	"github.com/MegaBytee/micro-fiber/auth"
	"github.com/MegaBytee/micro-fiber/cache"
	"github.com/MegaBytee/micro-fiber/limitter"
	"github.com/MegaBytee/micro-fiber/logger"
	"github.com/MegaBytee/micro-fiber/metrics"
	"github.com/MegaBytee/micro-fiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Service struct {
	config *Config
	fiber  *fiber.App
	auth   *auth.KeyAuth
	routes []*routes.ApiRoute
}

func NewService(config *Config) *Service {
	return &Service{
		config: config,
		fiber:  fiber.New(),
		auth:   auth.NewKeyAuth(config.AuthKeyLookup),
	}
}

func (s *Service) RegisterRoutes(routes []*routes.ApiRoute) {
	s.routes = routes
}

func (s *Service) loadRoutes() {
	log.Println("MicroFiber Service loadRoutes...")
	for _, route := range s.routes {
		route.Set(s.fiber)
	}
}
func (s *Service) Setup() {
	log.Println("MicroFiber Service Setup...")
	//default
	s.fiber.Use(helmet.New())
	s.fiber.Use(recover.New(
		recover.Config{
			EnableStackTrace: true,
		}))
	s.fiber.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	//setup logger
	if s.config.Logger {
		log.Println("MicroFiber Service Logger enabled...")
		logger.Setup(s.fiber)
	}
	//setup cache
	if s.config.Cache {
		log.Println("MicroFiber Service Cache enabled...")
		cache.Setup(s.fiber)
	}

	//setup limitter
	if s.config.Limitter {
		log.Println("MicroFiber Service Limitter enabled...")
		limitter.Setup(s.fiber)
	}

	//setup metrics
	if s.config.Metrics {
		log.Println("MicroFiber Service Metrics enabled...")
		metrics.Setup(s.fiber)
	}
	if s.config.ProtectedURLs != nil {
		log.Println("MicroFiber Service Auth enabled...")
		s.auth.Setup(s.fiber, s.config.ProtectedURLs)
		s.auth.ApiKeyLog()
	}

	s.fiber.Use(healthcheck.New())

	s.loadRoutes()
}

func (s *Service) Start() {

	log.Println("HTTP SERVER STARTED AT PORT:", s.config.Port)
	if err := s.fiber.Listen(":" + s.config.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
