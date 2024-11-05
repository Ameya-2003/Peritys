package main

import (
	"log"
	"wav-to-flac-converter/internal/ws"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// WebSocket route
	app.Post("/convert", ws.HandleConversion)

	log.Fatal(app.Listen(":3000"))
}
