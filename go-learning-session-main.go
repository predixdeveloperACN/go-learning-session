package main

import (
	"fmt"
	"github.com/predixdeveloperACN/go-learning-session/server"
	"github.com/predixdeveloperACN/go-learning-session/postgres"
)

func main() {
	fmt.Println("hello world")

	//init Postgres
	postgres.Open_postgres("user=postgres password=predix dbname=postgres sslmode=disable", "hospital")
	defer postgres.Database.Close()

	server.SetupServer()
	server.StartRestServer("8080")
}
