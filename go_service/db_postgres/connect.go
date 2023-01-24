package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func ConnectToDb() {
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://jb:12345678@localhost:54320/aot")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully.")
}

func GetDevices() {
	ConnectToDb()
	rows, err := db.Query(context.Background(), "SELECT * FROM apis_device	")
	if err != nil {
		panic(err)
		return
	}
	defer rows.Close()

	devices := make([]map[string]interface{}, 0)
	for rows.Next() {
		var id int
		var title string
		var description string
		var location_id string

		rows.Scan(&id, &title, &description, &location_id)
		devices = append(devices, map[string]interface{}{
			"id":          id,
			"title":       title,
			"description": description,
			"location_id": location_id,
		})
	}

	fmt.Println(devices)
	defer db.Close(context.Background())
}
