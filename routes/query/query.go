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
		"usersbyid": &graphql.Field{
			Type: types.UserType,
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data models.User
				id := p.Args["id_user"].(int)
				db.DB.First(&data, id)
				return data, nil
			},
		},
		"penjuals": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
		"penjualsbyid": &graphql.Field{
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
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Product

				// Ambil semua product + penjual
				err := db.DB.Preload("Penjual").Find(&data).Error
				if err != nil {
					return nil, err
				}

				var idUser uint = 0

				if val, ok := p.Args["id_user"]; ok && val != nil {
					idUser = uint(val.(int))
				}

				// Loop per product, cek apakah user memfavoritkan
				for i := range data {
					var fav models.Favorite
					err := db.DB.Where("id_product = ? AND id_user = ?", data[i].IDProduct, idUser).First(&fav).Error
					if err == nil {
						data[i].IDFavorite = &fav.IDFavorite
					} else {
						data[i].IDFavorite = nil
					}
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
				err := db.DB.Preload("Product").Preload("User").Find(&data).Error
				if err != nil {
					return nil, err
				}
				return data, nil
			},
		},
		"historys": &graphql.Field{
			Type: graphql.NewList(types.HistoryType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.History
				return data, db.DB.Find(&data).Error
			},
		},
		"keranjangs": &graphql.Field{
			Type: graphql.NewList(types.KeranjangType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Keranjang
				return data, db.DB.Find(&data).Error
			},
		},
		"idpenjuals": &graphql.Field{
			Type: graphql.NewList(types.PenjualType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var data []models.Penjual
				return data, db.DB.Find(&data).Error
			},
		},
		"alamats": &graphql.Field{
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
