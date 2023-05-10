package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "ubs"
)

type worker struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Department int    `json:"department"`
}

type logIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	//CONNECT TO DATABASE
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	//END OF CONNECTION

	app := fiber.New()
	app.Use(cors.New())

	ticket := 0

	var preferentialTickets []int
	var regularTickets []int

	//FUNÇOES DE ENVIO DE SENHA PARA BACK
	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Post("/preferential-ticket", func(context *fiber.Ctx) error {

		ticket = ticket + 1
		preferentialTickets = append(preferentialTickets, ticket)

		return context.JSON(ticket)
	})

	//ENVIA REQUISIÇÃO DE NOVO TICKET - REGULAR
	app.Post("/regular-ticket", func(context *fiber.Ctx) error {

		ticket = ticket + 1
		regularTickets = append(regularTickets, ticket)

		return context.JSON(ticket)
	})
	// =========================================================

	//FUNÇOES DE RESGATE DE SENHAS DO BACK
	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Get("/preferential-ticket", func(context *fiber.Ctx) error {
		return context.JSON(preferentialTickets)

	})

	//ENVIA REQUISIÇÃO DE NOVO TICKET - PREFERENCIAL
	app.Get("/regular-ticket", func(context *fiber.Ctx) error {
		return context.JSON(regularTickets)
	})
	// =========================================================

	//CRIA NOVO USUÁRIO
	app.Post("/new-user", func(context *fiber.Ctx) error {
		var newUser worker

		context.BodyParser(&newUser)

		sqlStatement := `
		INSERT INTO users (username, password, department)
		VALUES ($1, $2, $3)`

		_, err = db.Exec(sqlStatement, newUser.Username, newUser.Password, newUser.Department)
		if err != nil {
			return context.SendStatus(400)
		}

		return context.SendStatus(201)
	})

	//PROCURA USUÁRIO E CHECA SE A SENHA ESTA CORRETA
	app.Post("/log-in", func(context *fiber.Ctx) error {
		var newLogIn logIn
		var password string

		context.BodyParser(&newLogIn)

		row := db.QueryRow("SELECT password FROM users WHERE username= $1", newLogIn.Username)

		row.Scan(&password)
		if row == nil || password != newLogIn.Password {
			return context.SendStatus(400)
		}

		return context.SendStatus(200)
	})

	log.Fatal(app.Listen(":4000"))
}
