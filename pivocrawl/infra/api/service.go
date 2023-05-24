package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	app  *fiber.App
	port int
}

func NewService(port int) *Service {
	app := fiber.New()

	AddHealthRoutes(app)

	return &Service{
		app:  app,
		port: port,
	}
}

func (s *Service) GetAPI() *fiber.App {
	return s.app
}

func (s *Service) Start() error {
	address := fmt.Sprintf(":%d", s.port)
	log.Printf("Running API at address [%s]", address)
	return s.app.Listen(address)
}
