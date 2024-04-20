package main

import (
	"context"
	"encoding/json"
	"fmt"
	resolvers "github.com/Besufikad17/graphqldemo/resolvers"
	"github.com/graphql-go/graphql"
	"net/http"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    resolvers.QueryType,
	Mutation: resolvers.MutationType,
})

func executeQuery(query string, schema graphql.Schema, token string) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
		Context:       context.WithValue(context.Background(), "token", token),
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		result := executeQuery(r.URL.Query().Get("query"), Schema, tokenString)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8000", nil)
}
