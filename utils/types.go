package graphqldemo

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phoneNumber": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var MessageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AuthResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AuthResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
