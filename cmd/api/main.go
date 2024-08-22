package main

import (
	"back-sabervest/internal/models"
	"back-sabervest/internal/routes"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	var err error
	// global variable.
	models.DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/sabervest")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server running in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", routes.CreateRouter()))

}
