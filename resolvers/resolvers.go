package resolvers

import (
	db "github.com/Besufikad17/graphqldemo/db"
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var DB *gorm.DB = db.Init()

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"users": GetAllUsersResolver(DB),
		"user":  GetUserByIdResolver(DB),
	},
})

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"signup": SignUpResolver(DB),
		"login":  LoginResolver(DB),
		"add":    AddUserResolver(DB),
		"update": UpdateUserResolver(DB),
		"delete": DeleteUserResolver(DB),
	},
})
