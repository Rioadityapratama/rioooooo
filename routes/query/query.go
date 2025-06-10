package query

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type: graphql.NewList(types.UserType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.User
				return data, db.DB.Find(&data).Error
			},
		},
		"penjual": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
		"penjualbyid": &graphql.Field{
			Type: types.PenjualType,
			Args: graphql.FieldConfigArgument{
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data models.Penjual
				id := p.Args["id_penjual"].(int)
				db.DB.First(&data, id)
				return data, nil
			},
		},
		"products": &graphql.Field{
			Type: graphql.NewList(types.ProductType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Product
				err := db.DB.Preload("Penjual").Find(&data).Error
				if err != nil {
					return nil, err
				}
				return data, nil
			},
		},
		"checkouts": &graphql.Field{
			Type: graphql.NewList(types.CheckoutType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Checkout
				return data, db.DB.Find(&data).Error
			},
		},
		"favorites": &graphql.Field{
			Type: graphql.NewList(types.FavoriteType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Favorite
				return data, db.DB.Find(&data).Error
			},
		},
		"history": &graphql.Field{
			Type: graphql.NewList(types.HistoryType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.History
				return data, db.DB.Find(&data).Error
			},
		},
		"keranjang": &graphql.Field{
			Type: graphql.NewList(types.KeranjangType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Keranjang
				return data, db.DB.Find(&data).Error
			},
		},
		"idpenjual": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
		"alamat": &graphql.Field{
			Type: graphql.NewList(types.AlamatType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Alamat
				if err := db.DB.Preload("User").Find(&data).Error; err != nil {
					return nil, err
				}
				return data, nil
			},
		},
	},
})
