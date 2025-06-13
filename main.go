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
	//db.AutoMigrateTables()

	// Buat GraphQL schema (Query + Mutation)
	s := schema.NewSchema()

	// Setup handler GraphQL
	h := handler.New(&handler.Config{
		Schema:   &s,
		Pretty:   true,
		GraphiQL: true, // kalau mau GraphiQL nya diaktifkan
	})

	// Serve GraphQL API
	http.Handle("/graphql", h)

	// Serve static files (Frontend HTML kamu)
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Println("🚀 Server berjalan di: http://localhost:8080")
	log.Println("🚀 GraphQL: http://localhost:8080/graphql")
	//log.Println("🚀 Frontend: http://localhost:8080/register.html")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
