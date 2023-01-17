package main

import (
    "github.com/gofiber/fiber"
)

func main() {
    app := fiber.New()

    // Define GET route
    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

    // Define GET route with parameter
    app.Get("/users/:id", func(c *fiber.Ctx) {
        id := c.Params("id")
        c.JSON(fiber.Map{
            "id":   id,
            "name": "John Doe",
        })
    })

    // Define POST route
    app.Post("/users", func(c *fiber.Ctx) {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            c.Status(503).Send(err)
            return
        }
        c.JSON(user)
    })

    // Define PUT route
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

    // Define DELETE route
    app.Delete("/users/:id", func(c *fiber.Ctx) {
        id := c.Params("id")
        c.Send("User " + id + " has been deleted.")
    })

    // Start the server
    app.Listen(3000)
}

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
