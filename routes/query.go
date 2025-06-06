package routes

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.User
				return data, db.DB.Find(&data).Error
			},
		},
		"products": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Product
				return data, db.DB.Find(&data).Error
			},
		},
		"checkouts": &graphql.Field{
			Type: graphql.NewList(CheckoutType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Checkout
				return data, db.DB.Find(&data).Error
			},
		},
		"favorites": &graphql.Field{
			Type: graphql.NewList(FavoriteType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Favorite
				return data, db.DB.Find(&data).Error
			},
		},
		"history": &graphql.Field{
			Type: graphql.NewList(HistoryType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.History
				return data, db.DB.Find(&data).Error
			},
		},
		"keranjang": &graphql.Field{
			Type: graphql.NewList(KeranjangType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Keranjang
				return data, db.DB.Find(&data).Error
			},
		},
		"search": &graphql.Field{
			Type: graphql.NewList(SearchType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Search
				return data, db.DB.Find(&data).Error
			},
		},
		"penjual": &graphql.Field{
			Type: graphql.NewList(PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
	},
})
