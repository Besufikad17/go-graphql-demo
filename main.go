package main

import (
	"encoding/json"
	"fmt"
	db "github.com/Besufikad17/graphqldemo/db"
	utils "github.com/Besufikad17/graphqldemo/utils"
	"github.com/graphql-go/graphql"
	"net/http"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    utils.QueryType,
	Mutation: utils.MutationType,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	DB := db.Init()
	utils.NewResolver(DB)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), Schema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8000", nil)
}
