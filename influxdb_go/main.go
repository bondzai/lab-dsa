package main

import (
  	"fmt"

	"github.com/gofiber/fiber"
	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	app := fiber.New()
	client := influxdb2.NewClient("http://localhost:8086", "7mCH-C7vjw7ViMTI0hkDIfLm4fse5GA-2kX5BbGP-flkVtc-9sM6QKWQF9l7j7KnAMtSYtTZLlOmyZdOMWBqLQ==")
	fmt.Println(client)
	writeAPI := client.WriteAPIBlocking("jb", "demo")
	fmt.Println(writeAPI)

    app.Get("/", func(c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

	app.Post("/api/v1/devices/input", func(c *fiber.Ctx) {
		data := new(Data)
        if err := c.BodyParser(data); err != nil {
            c.Status(503).Send(err)
            return
        }
        fmt.Println(data)
        c.JSON(data)
    })

	app.Listen(4000)
}

type DataInflux struct {
	Timestamp string `json:"timestamp"`
	Measurement string `json:"measurement"`
	Air_port string `json:"air_port"`
	Location string `json:"location"`
	DevEUI string `json:"devEUI"`
	Device string `json:"device"`
	Rssi int `json:"rssi"`
	Occupancy *int `json:"occupancy"`
	Trigger_count *int `json:"trigger_count"`
	Install *int `json:"install"`
	Battery *float64 `json:"battery"`
}

type Data struct {
    ApplicationName string `json:"applicationName"`
    DeviceName string `json:"deviceName"`
    DevEUI string `json:"devEUI"`
    RxInfo []RxInfo `json:"rxInfo"`
    Object *Object `json:"object"`
    ObjectJSON *ObjectJSON `json:"objectJSON"`
}

type RxInfo struct {
	Rssi int `json:"rssi"`
}

type Object struct {
	Occupancy *int `json:"occupancy"`
	Trigger_count *int `json:"trigger_count"`
	Install *int `json:"install"`
	Battery *float64 `json:"battery"`
}

type ObjectJSON struct {
	Occupancy *int `json:"occupancy"`
	Trigger_count *int `json:"trigger_count"`
	Install *int `json:"install"`
	Battery *float64 `json:"battery"`
}
