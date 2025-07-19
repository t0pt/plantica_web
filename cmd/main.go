package main

import (
	"github.com/t0pt/plantica_web/cmd/events"
	"github.com/t0pt/plantica_web/cmd/server"
)

func main() {
	octopus := server.New("./templates")
	octopus.AddHandlers()
	events.CreateCalendarColumns()
	octopus.Listen(":8080")
}
