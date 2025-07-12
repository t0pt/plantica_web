package server

import (
	"github.com/gofiber/fiber/v2"

	"github.com/t0pt/plantica_web/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "github.com/t0pt/plantica_web",
			AppName:      "github.com/t0pt/plantica_web",
		}),

		db: database.New(),
	}

	return server
}
