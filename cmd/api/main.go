package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := mux.NewRouter()

	// dsn := os.Getenv("DSN")

	dsn := "postgres://myuser:mypasswd@localhost:5432/mydb"

	dep := InitDependencies(dsn)

	SetupRoutes(app, dep)

	log.Println("Server listening at :8080 ")
	return http.ListenAndServe(":8080", app)
}
