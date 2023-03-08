package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"test-architecture/adapter/api"
	"test-architecture/adapter/fixture"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configFile, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	scanner := bufio.NewScanner(configFile)
	scanner.Scan()
	username := scanner.Text()
	scanner.Scan()
	password := scanner.Text()
	scanner.Scan()
	host := scanner.Text()
	scanner.Scan()
	port := scanner.Text()
	scanner.Scan()
	database := scanner.Text()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := fixture.NewUserRepositoryDB(db)
	webserver := api.NewWebServer()
	webserver.Repository = repo
	webserver.Serve()
}
