// main.go
package main

import (
	"log"
	"net/http"

	"bakulos_grapghql/db"
	"bakulos_grapghql/routes/schema"

	"github.com/graphql-go/handler"
)

func main() {
	// Koneksi ke database
	db.ConnectDatabase()
	// db.AutoMigrateTables()

	// Buat GraphQL schema (Query + Mutation)
	schema := schema.NewSchema()

	// Setup handler dengan GraphiQL UI
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)
	log.Println("🚀 Server GraphQL berjalan di: http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
