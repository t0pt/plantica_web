package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{"Todos": Todos})
}

func AddTodoHandler(c *fiber.Ctx) error {
	title := c.FormValue("title")
	t := Todo{ID: len(Todos), Title: title}
	Todos = append(Todos, t)
	return c.Render("todo_row", t)
}

func ToggleTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i := range Todos {
		if Todos[i].ID == id {
			Todos[i].Done = !Todos[i].Done
			return c.Render("todo_row", Todos[i])
		}
	}
	return c.SendStatus(404)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	for i := range Todos {
		if Todos[i].ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return c.SendStatus(200)
		}
	}
	return c.SendStatus(404)
}
