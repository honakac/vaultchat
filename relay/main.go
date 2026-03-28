package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
)

const DEFAULT_ADDR string = ":8042"

func main() {
	addr := DEFAULT_ADDR
	if len(os.Args) == 2 {
		addr = os.Args[1]
	}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(addr, fiber.ListenConfig{
		EnablePrefork: false,
	}))
}
