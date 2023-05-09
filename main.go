package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	ticket := 0

	var preferentialTickets []int
	var regularTickets []int

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//FUNÇOES DE ENVIO DE SENHA PARA BACK
	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Post("/preferential-ticket", func(c *fiber.Ctx) error {

		ticket = ticket + 1
		preferentialTickets = append(preferentialTickets, ticket)

		return c.JSON(ticket)
	})

	//ENVIA REQUISIÇÃO DE NOVO TICKET - REGULAR
	app.Post("/regular-ticket", func(c *fiber.Ctx) error {

		ticket = ticket + 1
		regularTickets = append(regularTickets, ticket)

		return c.JSON(ticket)
	})
	// =========================================================

	//FUNÇOES DE RESGATE DE SENHAS DO BACK
	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Get("/preferential-ticket", func(c *fiber.Ctx) error {
		return c.JSON(preferentialTickets)

	})

	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Get("/regular-ticket", func(c *fiber.Ctx) error {
		return c.JSON(regularTickets)
	})
	// =========================================================

	log.Fatal(app.Listen(":4000"))
}
