package resolvers

import (
	"errors"
	handlers "github.com/Besufikad17/graphqldemo/handlers"
	models "github.com/Besufikad17/graphqldemo/models"
	utils "github.com/Besufikad17/graphqldemo/utils"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

func SignUpResolver(DB *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type:        utils.AuthResponseType,
		Description: "Admin signup",
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
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if p.Args["firstName"] == nil || p.Args["lastName"] == nil ||
				p.Args["email"] == nil || p.Args["phoneNumber"] == nil || p.Args["password"] == nil {
				return nil, errors.New("Please enter all fields!!")
			}

			user := &models.User{
				FirstName:   p.Args["firstName"].(string),
				LastName:    p.Args["lastName"].(string),
				Email:       p.Args["email"].(string),
				PhoneNumber: p.Args["phoneNumber"].(string),
				Password:    p.Args["password"].(string),
				Role:        models.Admin,
			}

			authHandler := handlers.NewAuthHandler(DB)
			token, err := authHandler.SignUp(user) // Pass the address of the user struct
			if err != nil {
				return nil, err
			}

			return models.AuthResponse{
				Message: "User Signed up successfully",
				Token:   token.(string),
			}, nil
		},
	}
}

func LoginResolver(DB *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type:        utils.AuthResponseType,
		Description: "Admin Login",
		Args: graphql.FieldConfigArgument{
			"loginText": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if p.Args["loginText"] == nil || p.Args["password"] == nil {
				return nil, errors.New("Please enter all fields!!")
			}

			authHandler := handlers.NewAuthHandler(DB)
			token, err := authHandler.Login(p.Args["loginText"].(string), p.Args["password"].(string)) // Pass the address of the user struct
			if err != nil {
				return nil, err
			}

			return models.AuthResponse{
				Message: "User Logged in successfully",
				Token:   token.(string),
			}, nil
		},
	}
}
