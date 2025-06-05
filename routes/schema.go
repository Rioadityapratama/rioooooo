package routes

import (
	"log"

	"github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	})
	if err != nil {
		log.Fatalf("❌ Gagal buat GraphQL schema: %v", err)
	}
	return schema
}
