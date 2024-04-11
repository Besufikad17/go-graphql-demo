package graphqldemo

import (
	"errors"

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
			Type:        graphql.NewList(UserType),
			Description: "Get All users",
			Args: graphql.FieldConfigArgument{
				"skip": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"take": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"text": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var skip int
				var take int
				var text string

				if p.Args["skip"] == nil {
					skip = 0
				} else {
					skip = p.Args["skip"].(int)
				}

				if p.Args["take"] == nil {
					take = 10
				} else {
					take = p.Args["take"].(int)
				}

				if p.Args["text"] == nil {
					text = ""
				} else {
					text = p.Args["text"].(string)
				}

				userHandler := handlers.NewUserHandler(DB)
				users, err := userHandler.GetAllUsers(&skip, &take, &text)
				if err != nil {
					return nil, err
				}
				return users, nil
			},
		},
		"user": &graphql.Field{
			Type:        UserType,
			Description: "Get user by Id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var id int

				if p.Args["id"] == nil {
					return nil, errors.New("Please enter id!!")
				} else {
					id = p.Args["id"].(int)
				}

				userHandler := handlers.NewUserHandler(DB)
				user, err := userHandler.GetUserById(id)

				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"add": &graphql.Field{
			Type:        UserType,
			Description: "Add user",
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
		"update": &graphql.Field{
			Type:        UserType,
			Description: "Update user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
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
				if p.Args["id"] == nil || p.Args["firstName"] == nil || p.Args["lastName"] == nil || p.Args["email"] == nil || p.Args["phoneNumber"] == nil {
					return nil, errors.New("Please enter all fields!!")
				}

				user := &models.User{
					FirstName:   p.Args["firstName"].(string),
					LastName:    p.Args["lastName"].(string),
					Email:       p.Args["email"].(string),
					PhoneNumber: p.Args["phoneNumber"].(string),
				}

				userHandler := handlers.NewUserHandler(DB)
				createdUser, err := userHandler.UpdateUser(uint(p.Args["id"].(int)), user) // Pass the address of the user struct
				if err != nil {
					return nil, err
				}

				return createdUser, nil
				return nil, nil
			},
		},
		"delete": &graphql.Field{
			Type:        MessageType,
			Description: "Delete user by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var id int

				if p.Args["id"] == nil {
					return nil, errors.New("Please enter id!!")
				} else {
					id = p.Args["id"].(int)
				}

				userHandler := handlers.NewUserHandler(DB)
				_, err := userHandler.DeleteUser(id)

				if err != nil {
					return nil, err
				}
				return models.Message{
					Text: "User Deleted successfully",
				}, nil
			},
		},
	},
})
