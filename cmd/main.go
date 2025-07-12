package main

import "github.com/t0pt/plantica_web/cmd/server"

func main() {
	octopus := server.New("./templates")
	octopus.AddHandlers()
	octopus.Listen(":8080")
}
