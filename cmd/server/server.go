package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type MainServer struct {
	*fiber.App
}

type Todo struct {
	ID    int
	Title string
	Done  bool
}

var Todos []Todo = []Todo{}

func New(location string) *MainServer {
	htmlEngine := html.New(location, ".html")
	return &MainServer{App: fiber.New(fiber.Config{Views: htmlEngine})}
}

func (s *MainServer) AddHandlers() {
	s.Get("/", IndexHandler)
	s.Post("/todos", AddTodoHandler)
	s.Put("/todos/:id", ToggleTodoHandler)
	s.Delete("/todos/:id", DeleteTodoHandler)
	s.Get("/calendar", CalendarHandler)
}
