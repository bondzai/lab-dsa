package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

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

func checkDeviceData() {
	fmt.Println("test")
}

func writeToInflux(dataInflux *DataInflux) {
	client := influxdb2.NewClient("http://localhost:8086", "7mCH-C7vjw7ViMTI0hkDIfLm4fse5GA-2kX5BbGP-flkVtc-9sM6QKWQF9l7j7KnAMtSYtTZLlOmyZdOMWBqLQ==")
	writeAPI := client.WriteAPI("jb", "demo")

	p := influxdb2.NewPointWithMeasurement(dataInflux.Measurement).
		AddField("rssi", dataInflux.Rssi).
		AddTag("air_port", dataInflux.Air_port).
		AddTag("devEUI", dataInflux.DevEUI).
		AddTag("location", dataInflux.Location).
		AddTag("device", dataInflux.Device).
		SetTime(dataInflux.Timestamp)

	if dataInflux.Occupancy != nil {
		p.AddField("occupancy", *dataInflux.Occupancy)
		fmt.Printf("value occ: %v\n", *dataInflux.Occupancy)
	}

	if dataInflux.Trigger_count != nil {
		p.AddField("trigger_count", *dataInflux.Trigger_count)
		fmt.Printf("value trig: %v\n", *dataInflux.Trigger_count)
	}

	if dataInflux.Install != nil {
		p.AddField("install", *dataInflux.Install)
		fmt.Printf("value install: %v\n", *dataInflux.Install)
	}

	if dataInflux.Battery != nil {
		p.AddField("battery", *dataInflux.Battery)
		fmt.Printf("value batt: %v\n", *dataInflux.Battery)
	}

	writeAPI.WritePoint(p)
	writeAPI.Flush()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		checkDeviceData()
		c.Send("JB")
	})

	app.Post("/api/v1/devices/input", func(c *fiber.Ctx) {

		data := new(Data)
		if err := c.BodyParser(data); err != nil {
			c.Status(503).Send(err)
			return
		}

		fmt.Println(" ")
		fmt.Println(" ******************* ")
		fmt.Println("data: ", data)
		fmt.Println(" ******************* ")

		// ========== data: common ==========
		dataInflux := new(DataInflux)
		fmt.Println("*********************************************************************")
		fmt.Printf("address dataInflux: %p\n", &dataInflux)
		fmt.Printf("address dataInflux: %p\n", dataInflux)
		fmt.Printf("address dataInflux: %v\n", *dataInflux)
		fmt.Println("*********************************************************************")

		dataInflux.Measurement = "PIR"
		dataInflux.Timestamp = time.Now()
		dataInflux.Air_port = "pending"
		dataInflux.Location = "pending"
		dataInflux.DevEUI = data.DevEUI
		dataInflux.Device = data.DeviceName
		dataInflux.Rssi = data.RxInfo[0].Rssi

		// ========== data: PIR / magnetic
		dataInflux.Occupancy = data.Object.Occupancy
		dataInflux.Trigger_count = data.Object.Trigger_count
		dataInflux.Install = data.Object.Install
		dataInflux.Battery = data.Object.Battery

		fmt.Println(dataInflux.Occupancy)
		fmt.Println(dataInflux.Trigger_count)
		fmt.Println(dataInflux.Install)
		fmt.Println(dataInflux.Battery)

		// monitor data for debug
		fmt.Println("***** dataInflux: start *****")
		fmt.Println("  	data influx timestamp: ", dataInflux.Timestamp)
		fmt.Println("	data influx measurement: ", dataInflux.Measurement)
		fmt.Println("	data influx airport: (get from postgresql) ", dataInflux.Air_port)
		fmt.Println("	data influx location: (get from postgresql)", dataInflux.Location)
		fmt.Println("	data influx devEUI: ", dataInflux.DevEUI)
		fmt.Println("	data influx device: ", dataInflux.Device)
		fmt.Println("	data influx rssi: ", dataInflux.Rssi)
		fmt.Println("	data influx occupancy: ", dataInflux.Occupancy)
		fmt.Println("	data influx trigger_count: ", dataInflux.Trigger_count)
		fmt.Println("	data influx install: ", dataInflux.Install)
		fmt.Println("	data influx battery: ", dataInflux.Battery)
		fmt.Println("***** dataInflux: end *****")

		writeToInflux(dataInflux)
		c.JSON(data)
	})

	app.Listen(4000)
}
