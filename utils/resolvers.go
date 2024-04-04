package graphqldemo

import (
	handlers "github.com/Besufikad17/graphqldemo/handlers"
	models "github.com/Besufikad17/graphqldemo/models"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewResolver(db *gorm.DB) {
	DB = db
}

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
					if int(Users[i].ID) == id {
						return Users[i], nil
					}
				}
				return nil, nil
			},
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"add": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"firstName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"lastName": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"phoneNumber": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := &models.User{
					FirstName:   p.Args["firstName"].(string),
					LastName:    p.Args["lastName"].(string),
					Email:       p.Args["email"].(string),
					PhoneNumber: p.Args["phoneNumber"].(string),
				}

				userHandler := handlers.NewUserHandler(DB)
				createdUser, err := userHandler.AddUser(user) // Pass the address of the user struct
				if err != nil {
					return nil, err
				}

				return createdUser, nil
			},
		},
	},
})
