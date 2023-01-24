package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	app := fiber.New()
	client := influxdb2.NewClient("http://localhost:8086", "7mCH-C7vjw7ViMTI0hkDIfLm4fse5GA-2kX5BbGP-flkVtc-9sM6QKWQF9l7j7KnAMtSYtTZLlOmyZdOMWBqLQ==")
	writeAPI := client.WriteAPI("jb", "demo")

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("JB")
	})

	app.Post("/api/v1/devices/input", func(c *fiber.Ctx) {

		// get rawdata & convert from JSON to go Struct
		data := new(Data)
		if err := c.BodyParser(data); err != nil {
			c.Status(503).Send(err)
			return
		}

		// setup data before saving into InfluxDB
		// ========== data: static ==========
		dataInflux := new(DataInflux)
		dataInflux.Timestamp = time.Now()
		dataInflux.Measurement = "1"
		dataInflux.Air_port = "pending"
		dataInflux.Location = "pending"
		dataInflux.DevEUI = data.DevEUI
		dataInflux.Device = data.DeviceName
		dataInflux.Rssi = data.RxInfo[0].Rssi
		// ========== data: dynamic ==========
		// ========== data: PIR
		dataInflux.Occupancy = data.Object.Occupancy
		dataInflux.Trigger_count = data.Object.Trigger_count
		// ========== data: magnetic

		// monitor data for debug
		fmt.Println(" ===== dataInflux: start =====")
		fmt.Println("data influx: ", dataInflux)
		fmt.Println(" ===== dataInflux: detail =====")
		fmt.Println("data influx timestamp: ", dataInflux.Timestamp)
		fmt.Println("data influx measurement: ", dataInflux.Measurement)
		fmt.Println("data influx airport: (get from postgresql) ", dataInflux.Air_port)
		fmt.Println("data influx location: (get from postgresql)", dataInflux.Location)
		fmt.Println("data influx devEUI: ", dataInflux.DevEUI)
		fmt.Println("data influx device: ", dataInflux.Device)
		fmt.Println("data influx rssi: ", dataInflux.Rssi)
		fmt.Println(" ===== dataInflux: end =====")

		p := influxdb2.NewPointWithMeasurement("JB3").
			AddTag("unit", "unit of measure").
			AddField("measurement", 22.2).
			AddField("air_port", dataInflux.Air_port).
			AddField("location", dataInflux.Location).
			AddField("devEUI", dataInflux.DevEUI).
			AddField("device", dataInflux.Device).
			AddField("rssi", dataInflux.Rssi).
			SetTime(dataInflux.Timestamp)
		writeAPI.WritePoint(p)
		writeAPI.Flush()

		fmt.Println(data)
		c.JSON(data)
	})

	app.Listen(4000)
}

type DataInflux struct {
	Timestamp     time.Time `json:"timestamp"`
	Measurement   string    `json:"measurement"`
	Air_port      string    `json:"air_port"`
	Location      string    `json:"location"`
	DevEUI        string    `json:"devEUI"`
	Device        string    `json:"device"`
	Rssi          int       `json:"rssi"`
	Occupancy     *int      `json:"occupancy"`
	Trigger_count *int      `json:"trigger_count"`
	Install       *int      `json:"install"`
	Battery       *float64  `json:"battery"`
}

type Data struct {
	ApplicationName string      `json:"applicationName"`
	DeviceName      string      `json:"deviceName"`
	DevEUI          string      `json:"devEUI"`
	RxInfo          []RxInfo    `json:"rxInfo"`
	Object          *Object     `json:"object"`
	ObjectJSON      *ObjectJSON `json:"objectJSON"`
}

type RxInfo struct {
	Rssi int `json:"rssi"`
}

type Object struct {
	Occupancy     *int     `json:"occupancy"`
	Trigger_count *int     `json:"trigger_count"`
	Install       *int     `json:"install"`
	Battery       *float64 `json:"battery"`
}

type ObjectJSON struct {
	Occupancy     *int     `json:"occupancy"`
	Trigger_count *int     `json:"trigger_count"`
	Install       *int     `json:"install"`
	Battery       *float64 `json:"battery"`
}