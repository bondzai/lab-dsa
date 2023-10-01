package main

import (
    "fmt"

    "github.com/gofiber/fiber"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

    app.Get("/users/:id", func(c *fiber.Ctx) {
        id := c.Params("id")
        c.JSON(fiber.Map{
            "id":   id,
            "name": "John Doe",
        })
    })

    app.Post("/users", func(c *fiber.Ctx) {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            c.Status(503).Send(err)
            return
        }
        fmt.Println(user)
        c.JSON(user)
    })

    app.Put("/users/:id", func(c *fiber.Ctx) {
        id := c.Params("id")
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            c.Status(503).Send(err)
            return
        }
        c.JSON(fiber.Map{
            "id":   id,
            "name": user.Name,
        })
    })

    app.Delete("/users/:id", func(c *fiber.Ctx) {
        id := c.Params("id")
        c.Send("User " + id + " has been deleted.")
    })

    app.Listen(4000)
}

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
