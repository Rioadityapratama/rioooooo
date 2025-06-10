package schema

import (
	"bakulos_grapghql/routes/mutation"
	"bakulos_grapghql/routes/query"
	"log"

	"github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.RootQuery,
		Mutation: mutation.RootMutation,
	})
	if err != nil {
		log.Fatalf("❌ Gagal buat GraphQL schema: %v", err)
	}
	return schema
}
