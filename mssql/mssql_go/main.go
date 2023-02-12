package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// Connect to the database server
	connString := "server=localhost;user id=sa;password=12345678;port=1433;database=master"
	db, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
		return
	}
	defer db.Close()

	// Make a test database connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err.Error())
		return
	}
	fmt.Println("Successfully connected to SQL server!")
}
