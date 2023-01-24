package db_postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func ConnectToPostgres() {
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://jb:12345678@localhost:54320/aot")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully.")
}
