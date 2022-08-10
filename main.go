package main

import (
	"log"

	// "net/http"
	// _ "net/http/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	app := fiber.New()

	app.Use(pprof.New())

	app.Get("/", handleRequest)
	log.Fatal(app.Listen(":8080"))
}

func handleRequest(c *fiber.Ctx) error {
	log.Printf("handling request from: %s", c.Context().RemoteAddr())
	if _, err := c.WriteString(c.Context().RemoteAddr().String()); err != nil {
		log.Printf("could not write IP: %s", err)
	}
	return nil
}
