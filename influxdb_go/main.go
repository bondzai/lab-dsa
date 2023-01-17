package main

import (
  	"fmt"
	"github.com/gofiber/fiber"
   	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	app := fiber.New()
	client := influxdb2.NewClient("http://localhost:8086", "7mCH-C7vjw7ViMTI0hkDIfLm4fse5GA-2kX5BbGP-flkVtc-9sM6QKWQF9l7j7KnAMtSYtTZLlOmyZdOMWBqLQ==")
  	writeAPI := client.WriteAPI("jb", "demo")

    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

	app.Post("/api/v1/devices/input", func(c *fiber.Ctx) {
		writeAPI.WriteRecord(fmt.Sprintf("stat3,unit=temperature avg=%f", 45.0))
		writeAPI.Flush()
    })

	defer client.Close()
	app.Listen(4000)
}
