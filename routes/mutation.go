package routes

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"github.com/graphql-go/graphql"
)

var DB = db.DB

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		// USER
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"nama":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user := models.User{
					Nama:     p.Args["nama"].(string),
					Email:    p.Args["email"].(string),
					Password: p.Args["password"].(string),
				}
				if err := DB.Create(&user).Error; err != nil {
					return nil, err
				}
				return user, nil
			},
		},

		"updateUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"id_user":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"nama":     &graphql.ArgumentConfig{Type: graphql.String},
				"email":    &graphql.ArgumentConfig{Type: graphql.String},
				"password": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				id := p.Args["id_user"].(int)
				if err := DB.First(&user, id).Error; err != nil {
					return nil, err
				}
				if v, ok := p.Args["nama"].(string); ok {
					user.Nama = v
				}
				if v, ok := p.Args["email"].(string); ok {
					user.Email = v
				}
				if v, ok := p.Args["password"].(string); ok {
					user.Password = v
				}
				if err := DB.Save(&user).Error; err != nil {
					return nil, err
				}
				return user, nil
			},
		},

		"deleteUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id_user": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if err := DB.Delete(&models.User{}, p.Args["id_user"].(int)).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},

		// PRODUCT
		"createProduct": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"kategori":   &graphql.ArgumentConfig{Type: graphql.String},
				"size":       &graphql.ArgumentConfig{Type: graphql.String},
				"deskripsi":  &graphql.ArgumentConfig{Type: graphql.String},
				"brand":      &graphql.ArgumentConfig{Type: graphql.String},
				"price":      &graphql.ArgumentConfig{Type: graphql.Float},
				"image":      &graphql.ArgumentConfig{Type: graphql.String},
				"warna":      &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := models.Product{
					IDPenjual: uint(p.Args["id_penjual"].(int)),
					Kategori:  p.Args["kategori"].(string),
					Size:      p.Args["size"].(string),
					Deskripsi: p.Args["deskripsi"].(string),
					Brand:     p.Args["brand"].(string),
					Price:     int(p.Args["price"].(float64)),
					Image:     p.Args["image"].(string),
					Warna:     p.Args["warna"].(string),
				}
				if err := DB.Create(&product).Error; err != nil {
					return nil, err
				}
				return product, nil
			},
		},

		"deleteProduct": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if err := DB.Delete(&models.Product{}, p.Args["id_product"].(int)).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},

		// KERANJANG
		"createKeranjang": &graphql.Field{
			Type: KeranjangType,
			Args: graphql.FieldConfigArgument{
				"id_product": &graphql.ArgumentConfig{Type: graphql.Int},
				"id_user":    &graphql.ArgumentConfig{Type: graphql.Int},
				"jumlah":     &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				keranjang := models.Keranjang{
					IDProduct: uint(p.Args["id_product"].(int)),
					IDUser:    uint(p.Args["id_user"].(int)),
					Jumlah:    p.Args["jumlah"].(int),
				}
				if err := DB.Create(&keranjang).Error; err != nil {
					return nil, err
				}
				return keranjang, nil
			},
		},

		"deleteKeranjang": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id_keranjang": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if err := DB.Delete(&models.Keranjang{}, p.Args["id_keranjang"].(int)).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},
	},
})
