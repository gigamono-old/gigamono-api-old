package main

import (
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %v 👋🏾!", c.Params("name"))
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))
}
