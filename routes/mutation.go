package routes

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

// Helper functions to get arguments from ResolveParams
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

// Root Mutation
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{

		// ==================== USER ====================
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
				if _, exists := p.Args["email"]; exists {
					return nil, fmt.Errorf("field 'email' tidak dapat diubah")
				}
				id := getInt(p, "id_user")
				var user models.User
				if err := db.DB.First(&user, id).Error; err != nil {
					return nil, fmt.Errorf("user dengan id %d tidak ditemukan", id)
				}
				updates := map[string]interface{}{}
				if v := getString(p, "nama"); v != "" {
					updates["nama"] = v
				}
				if v := getString(p, "password"); v != "" {
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
					if err != nil {
						return nil, fmt.Errorf("gagal meng-hash password: %v", err)
					}
					updates["password"] = string(hashedPassword)
				}
				if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
					return nil, fmt.Errorf("gagal update data user: %v", err)
				}
				db.DB.First(&user, id)
				return user, nil
			},
		},

		// ==================== PENJUAL ====================
		"createPenjual": &graphql.Field{
			Type: PenjualType,
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
				penjual := models.Penjual{
					Nama:     p.Args["nama"].(string),
					Email:    p.Args["email"].(string),
					Password: string(hashedPassword),
				}
				if err := db.DB.Create(&penjual).Error; err != nil {
					return nil, err
				}
				return penjual, nil
			},
		},

		"updatePenjual": &graphql.Field{
			Type: PenjualType,
			Args: graphql.FieldConfigArgument{
				"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"nama":       &graphql.ArgumentConfig{Type: graphql.String},
				"password":   &graphql.ArgumentConfig{Type: graphql.String},
				"email":      &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if _, exists := p.Args["email"]; exists {
					return nil, fmt.Errorf("field 'email' tidak dapat diubah")
				}
				id := getInt(p, "id_penjual")
				var penjual models.Penjual
				if err := db.DB.First(&penjual, id).Error; err != nil {
					return nil, fmt.Errorf("penjual dengan id %d tidak ditemukan", id)
				}
				updates := map[string]interface{}{}
				if v := getString(p, "nama"); v != "" {
					updates["nama"] = v
				}
				if v := getString(p, "password"); v != "" {
					hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
					if err != nil {
						return nil, fmt.Errorf("gagal meng-hash password: %v", err)
					}
					updates["password"] = string(hashedPassword)
				}
				if err := db.DB.Model(&penjual).Updates(updates).Error; err != nil {
					return nil, fmt.Errorf("gagal update data penjual: %v", err)
				}
				db.DB.First(&penjual, id)
				return penjual, nil
			},
		},

		// ==================== PRODUCT ====================
		"createProduct": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id_penjual":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"namaproduk":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"kategori":    &graphql.ArgumentConfig{Type: graphql.String},
				"size":        &graphql.ArgumentConfig{Type: graphql.String},
				"deskripsi":   &graphql.ArgumentConfig{Type: graphql.String},
				"brand":       &graphql.ArgumentConfig{Type: graphql.String},
				"price":       &graphql.ArgumentConfig{Type: graphql.Float},
				"image":       &graphql.ArgumentConfig{Type: graphql.String},
				"warna":       &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idPenjual := getInt(p, "id_penjual")
				var penjual models.Penjual
				if err := db.DB.First(&penjual, idPenjual).Error; err != nil {
					return nil, fmt.Errorf("penjual dengan id %d tidak ditemukan", idPenjual)
				}
				price := 0
				if v, ok := p.Args["price"].(float64); ok {
					price = int(v)
				}
				product := models.Product{
					IDPenjual:  uint(idPenjual),
					NamaProduk: getString(p, "namaproduk"),
					Kategori:   getString(p, "kategori"),
					Size:       getString(p, "size"),
					Deskripsi:  getString(p, "deskripsi"),
					Brand:      getString(p, "brand"),
					Price:      price,
					Image:      getString(p, "image"),
					Warna:      getString(p, "warna"),
				}
				if err := db.DB.Create(&product).Error; err != nil {
					return nil, fmt.Errorf("gagal menyimpan produk: %v", err)
				}
				db.DB.Preload("Penjual").First(&product, product.IDProduct)
				return product, nil
			},
		},

		"updateProduct": &graphql.Field{
			Type: ProductType,
			Args: graphql.FieldConfigArgument{
				"id_product":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"namaproduk":  &graphql.ArgumentConfig{Type: graphql.String},
				"size":        &graphql.ArgumentConfig{Type: graphql.String},
				"brand":       &graphql.ArgumentConfig{Type: graphql.String},
				"deskripsi":   &graphql.ArgumentConfig{Type: graphql.String},
				"kategori":    &graphql.ArgumentConfig{Type: graphql.String},
				"price":       &graphql.ArgumentConfig{Type: graphql.Float},
				"image":       &graphql.ArgumentConfig{Type: graphql.String},
				"warna":       &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := getInt(p, "id_product")
				var product models.Product
				if err := db.DB.First(&product, id).Error; err != nil {
					return nil, fmt.Errorf("produk tidak ditemukan")
				}
				updates := map[string]interface{}{}
				if v := getString(p, "namaproduk"); v != "" {
					updates["nama"] = v
				}
				if v := getString(p, "size"); v != "" {
					updates["size"] = v
				}
				if v := getString(p, "brand"); v != "" {
					updates["brand"] = v
				}
				if v := getString(p, "deskripsi"); v != "" {
					updates["deskripsi"] = v
				}
				if v := getString(p, "kategori"); v != "" {
					updates["kategori"] = v
				}
				if v := getString(p, "image"); v != "" {
					updates["image"] = v
				}
				if v := getString(p, "warna"); v != "" {
					updates["warna"] = v
				}
				if v, ok := p.Args["price"].(float64); ok {
					updates["price"] = int(v)
				}
				if err := db.DB.Model(&product).Updates(updates).Error; err != nil {
					return nil, fmt.Errorf("gagal update")
				}
				db.DB.Preload("Penjual").First(&product, id)
				return product, nil
			},
		},

		"deleteProduct": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "DeleteProductResponse",
				Fields: graphql.Fields{
					"message": &graphql.Field{Type: graphql.String},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := getInt(p, "id_product")
				var product models.Product
				if err := db.DB.First(&product, id).Error; err != nil {
					return map[string]interface{}{"message": "Produk tidak ditemukan"}, nil
				}
				if err := db.DB.Delete(&product).Error; err != nil {
					return map[string]interface{}{"message": "Gagal menghapus produk"}, nil
				}
				return map[string]interface{}{"message": "Produk berhasil dihapus"}, nil
			},
		},

		// ==================== KERANJANG ====================
		"createKeranjang": &graphql.Field{
			Type: KeranjangType,
			Args: graphql.FieldConfigArgument{
				"id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"id_user":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"jumlah":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idProduct := getInt(p, "id_product")
				idUser := getInt(p, "id_user")
				jumlah := getInt(p, "jumlah")

				var user models.User
				if err := db.DB.First(&user, idUser).Error; err != nil {
					return nil, fmt.Errorf("user tidak ditemukan")
				}
				var product models.Product
				if err := db.DB.First(&product, idProduct).Error; err != nil {
					return nil, fmt.Errorf("produk tidak ditemukan")
				}
				keranjang := models.Keranjang{
					IDProduct: uint(idProduct),
					IDUser:    uint(idUser),
					Jumlah:    jumlah,
				}
				if err := db.DB.Create(&keranjang).Error; err != nil {
					return nil, err
				}
				db.DB.Preload("User").Preload("Product").First(&keranjang, keranjang.IDKeranjang)
				return keranjang, nil
			},
		},

		"updateKeranjang": &graphql.Field{
			Type: KeranjangType,
			Args: graphql.FieldConfigArgument{
				"id_keranjang": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"id_product":   &graphql.ArgumentConfig{Type: graphql.Int},
				"jumlah":       &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				idKeranjang := getInt(p, "id_keranjang")
				var keranjang models.Keranjang
				if err := db.DB.First(&keranjang, idKeranjang).Error; err != nil {
					return nil, fmt.Errorf("keranjang dengan id %d tidak ditemukan", idKeranjang)
				}

				updates := map[string]interface{}{}

				if v := getInt(p, "id_product"); v != 0 {
					var product models.Product
					if err := db.DB.First(&product, v).Error; err != nil {
						return nil, fmt.Errorf("produk dengan id %d tidak ditemukan", v)
					}
					updates["id_product"] = v
				}
				if v := getInt(p, "jumlah"); v != 0 {
					updates["jumlah"] = v
				}

				if len(updates) == 0 {
					return nil, fmt.Errorf("tidak ada field yang diperbarui")
				}

				if err := db.DB.Model(&keranjang).Updates(updates).Error; err != nil {
					return nil, fmt.Errorf("gagal mengupdate keranjang: %v", err)
				}

				if err := db.DB.Preload("User").Preload("Product").First(&keranjang, idKeranjang).Error; err != nil {
					return nil, fmt.Errorf("gagal mengambil ulang keranjang")
				}
				return keranjang, nil
			},
		},

		"deleteKeranjang": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "DeleteKeranjangResponse",
				Fields: graphql.Fields{
					"message": &graphql.Field{Type: graphql.String},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"id_keranjang": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := getInt(p, "id_keranjang")
				var keranjang models.Keranjang
				if err := db.DB.First(&keranjang, id).Error; err != nil {
					return map[string]interface{}{"message": "Keranjang tidak ditemukan"}, nil
				}
				if err := db.DB.Delete(&keranjang).Error; err != nil {
					return map[string]interface{}{"message": "Gagal menghapus keranjang"}, nil
				}
				return map[string]interface{}{"message": "Keranjang berhasil dihapus"}, nil
			},
		},
	},
})
