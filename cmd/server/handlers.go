package server

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t0pt/plantica_web/cmd/events"
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

func CalendarHandler(c *fiber.Ctx) error {
	var hours []int
	h := 5
	for {
		hours = append(hours, h)
		if h == 4 {
			break
		}
		h = (h + 1) % 24
	}
	return c.Render("calendar", fiber.Map{
		"CalendarColumns": events.CalendarColumns,
		"TimelineHours":   hours,
	})
}

func TaskCompleteHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}
	for _, task := range events.Tasks {
		if task.ID == int64(id) {
			task.Done = !task.Done
			// Render the updated task using the partial template
			return c.Render("task", task, "task")
		}
	}
	return c.SendStatus(404)
}
