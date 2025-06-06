package routes

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

func getString(p graphql.ResolveParams, key string) string {
	if val, ok := p.Args[key].(string); ok {
		return val
	}
	return ""
}

func getInt(p graphql.ResolveParams, key string) int {
	if val, ok := p.Args[key].(int); ok {
		return val
	}
	return 0
}

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"nama":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Args["password"].(string)), bcrypt.DefaultCost)
				if err != nil {
					return nil, err
				}
				user := models.User{
					Nama:     p.Args["nama"].(string),
					Email:    p.Args["email"].(string),
					Password: string(hashedPassword),
				}
				if err := db.DB.Create(&user).Error; err != nil {
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
				id := getInt(p, "id_user")
				if err := db.DB.First(&user, id).Error; err != nil {
					return nil, err
				}
				if v := getString(p, "nama"); v != "" {
					user.Nama = v
				}
				if v := getString(p, "email"); v != "" {
					user.Email = v
				}
				if v := getString(p, "password"); v != "" {
					hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
					user.Password = string(hashedPassword)
				}
				if err := db.DB.Save(&user).Error; err != nil {
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
				if err := db.DB.Delete(&models.User{}, getInt(p, "id_user")).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},

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
				price := 0
				if v, ok := p.Args["price"].(float64); ok {
					price = int(v)
				}
				product := models.Product{
					IDPenjual: uint(getInt(p, "id_penjual")),
					Kategori:  getString(p, "kategori"),
					Size:      getString(p, "size"),
					Deskripsi: getString(p, "deskripsi"),
					Brand:     getString(p, "brand"),
					Price:     price,
					Image:     getString(p, "image"),
					Warna:     getString(p, "warna"),
				}
				if err := db.DB.Create(&product).Error; err != nil {
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
				if err := db.DB.Delete(&models.Product{}, getInt(p, "id_product")).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},

		"createKeranjang": &graphql.Field{
			Type: KeranjangType,
			Args: graphql.FieldConfigArgument{
				"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"id_user":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"jumlah":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				keranjang := models.Keranjang{
					IDProduct: uint(getInt(p, "id_product")),
					IDUser:    uint(getInt(p, "id_user")),
					Jumlah:    getInt(p, "jumlah"),
				}
				if err := db.DB.Create(&keranjang).Error; err != nil {
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
				if err := db.DB.Delete(&models.Keranjang{}, getInt(p, "id_keranjang")).Error; err != nil {
					return false, err
				}
				return true, nil
			},
		},
	},
})
