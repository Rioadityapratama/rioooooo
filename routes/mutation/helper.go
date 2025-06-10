package mutation

import "github.com/graphql-go/graphql"

func getString(p graphql.ResolveParams, key string) string {
	if val, ok := p.Args[key].(string); ok {
		return val
	}
	return ""
}
func getInt(p graphql.ResolveParams, key string) int {
	if val, ok := p.Args[key].(int); ok {
		return val
	}
	return 0
}
