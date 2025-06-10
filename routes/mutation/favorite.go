package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"
	"github.com/graphql-go/graphql"
)

var CreateFavorite = &graphql.Field{
	Type: types.FavoriteType,
	Args: graphql.FieldConfigArgument{
		"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_user":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idProduct := getInt(p, "id_product")
		idUser := getInt(p, "id_user")
		var user models.User
		if err := db.DB.First(&user, idUser).Error; err != nil {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
		var product models.Product
		if err := db.DB.First(&product, idProduct).Error; err != nil {
			return nil, fmt.Errorf("produk tidak ditemukan")
		}
		favorite := models.Favorite{
			IDProduct: uint(idProduct),
			IDUser:    uint(idUser),
		}
		if err := db.DB.Create(&favorite).Error; err != nil {
			return nil, fmt.Errorf("gagal menyimpan favorite: %v", err)
		}
		db.DB.Preload("User").Preload("Product").First(&favorite, favorite.IDFavorite)
		return favorite, nil
	},
}

var UpdateFavorite = &graphql.Field{
	Type: types.FavoriteType,
	Args: graphql.FieldConfigArgument{
		"id_favorite": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_product":  &graphql.ArgumentConfig{Type: graphql.Int},
		"id_user":     &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idFavorite := getInt(p, "id_favorite")
		var favorite models.Favorite
		if err := db.DB.First(&favorite, idFavorite).Error; err != nil {
			return nil, fmt.Errorf("favorite dengan id %d tidak ditemukan", idFavorite)
		}
		updates := map[string]interface{}{}
		if v := getInt(p, "id_product"); v != 0 {
			var product models.Product
			if err := db.DB.First(&product, v).Error; err != nil {
				return nil, fmt.Errorf("produk dengan id %d tidak ditemukan", v)
			}
			updates["id_product"] = v
		}
		if v := getInt(p, "id_user"); v != 0 {
			var user models.User
			if err := db.DB.First(&user, v).Error; err != nil {
				return nil, fmt.Errorf("user dengan id %d tidak ditemukan", v)
			}
			updates["id_user"] = v
		}
		if len(updates) == 0 {
			return nil, fmt.Errorf("tidak ada field yang diperbarui")
		}
		if err := db.DB.Model(&favorite).Updates(updates).Error; err != nil {
			return nil, fmt.Errorf("gagal mengupdate favorite: %v", err)
		}
		if err := db.DB.Preload("User").Preload("Product").First(&favorite, idFavorite).Error; err != nil {
			return nil, fmt.Errorf("gagal mengambil ulang favorite")
		}
		return favorite, nil
	},
}

var DeleteFavorite = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "DeleteFavoriteResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"id_favorite": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_favorite")
		var favorite models.Favorite
		if err := db.DB.First(&favorite, id).Error; err != nil {
			return map[string]interface{}{"message": "Favorite tidak ditemukan"}, nil
		}
		if err := db.DB.Delete(&favorite).Error; err != nil {
			return map[string]interface{}{"message": "Gagal menghapus favorite"}, nil
		}
		return map[string]interface{}{"message": "Favorite berhasil dihapus"}, nil
	},
}
