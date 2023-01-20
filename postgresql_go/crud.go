package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func main() {
	fmt.Println(" ===== start =====")
	app := fiber.New()
	app.Use(ConnectToDb())

	app.Get("/users", GetUsers)
	app.Post("/users", CreateUser)
	app.Put("/users/:id", UpdateUser)
	app.Delete("/users/:id", DeleteUser)

	app.Listen(3000)
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

func GetUsers(c *fiber.Ctx) {
	rows, err := db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		c.Send(err)
		return
	}
	defer rows.Close()

	users := make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var name string
		var email string
		var password string

		rows.Scan(&id, &name, &email, &password)
		users = append(users, map[string]interface{}{
			"id":       id,
			"name":     name,
			"email":    email,
			"password": password,
		})
	}

	c.JSON(users)
}

func CreateUser(c *fiber.Ctx) {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := db.Exec(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, password)
	if err != nil {
		c.Send(err)
		return
	}

	c.Send("User created successfully!")
}

func UpdateUser(c *fiber.Ctx) {
	id := c.Params("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := db.Exec(context.Background(), "UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4", name, email, password, id)
	if err != nil {
		c.Send(err)
		return
	}

	c.Send("User updated successfully!")
}

func DeleteUser(c *fiber.Ctx) {
	id := c.Params("id")

	_, err := db.Exec(context.Background(), "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.Send(err)
		return
	}

	c.Send("User deleted successfully!")
}
