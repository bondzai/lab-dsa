package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func main() {
	app := fiber.New()
	app.Use(ConnectToDb())

	app.Get("/", func(c *fiber.Ctx) {
		_, err := db.Exec(context.Background(), "SELECT 1")
		if err != nil {
			c.Send(err)
			return
		}
		fmt.Println("===")
		c.Send("Hello, World!")
	})

	app.Listen(3001)
}

func ConnectToDb() fiber.Handler {
	return func(c *fiber.Ctx) {
		var err error
		db, err = pgx.Connect(context.Background(), "postgres://postgres:12345678@localhost:5432/go")
		if err != nil {
			c.Send(err)
			return
		}
		defer db.Close(context.Background())

		c.Next()
	}
}
