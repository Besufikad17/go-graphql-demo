package graphqldemo

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return Users, nil
			},
		},
		"user": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				for i := 0; i < len(Users); i++ {
					if Users[i].ID == id {
						return Users[i], nil
					}
				}
				return nil, nil
			},
		},
	},
})
