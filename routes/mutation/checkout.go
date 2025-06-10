package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
)

var CreateCheckout = &graphql.Field{
	Type: types.CheckoutType,
	Args: graphql.FieldConfigArgument{
		"id_user":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_product":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_keranjang":      &graphql.ArgumentConfig{Type: graphql.Int}, // optional
		"id_alamat":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"metode_pengiriman": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"pembayaran":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"jumlah":            &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idUser := getInt(p, "id_user")
		idProduct := getInt(p, "id_product")
		idAlamat := getInt(p, "id_alamat")
		var user models.User
		if err := db.DB.First(&user, idUser).Error; err != nil {
			return nil, fmt.Errorf("user dengan id %d tidak ditemukan", idUser)
		}
		var product models.Product
		if err := db.DB.First(&product, idProduct).Error; err != nil {
			return nil, fmt.Errorf("produk dengan id %d tidak ditemukan", idProduct)
		}
		var alamat models.Alamat
		if err := db.DB.First(&alamat, idAlamat).Error; err != nil {
			return nil, fmt.Errorf("alamat dengan id %d tidak ditemukan", idAlamat)
		}
		var idKeranjang *uint = nil
		if val, ok := p.Args["id_keranjang"]; ok {
			tempID := uint(val.(int))
			var keranjang models.Keranjang
			if err := db.DB.First(&keranjang, tempID).Error; err != nil {
				return nil, fmt.Errorf("keranjang dengan id %d tidak ditemukan", tempID)
			}
			idKeranjang = &tempID
		}
		checkout := models.Checkout{
			IDUser:           uint(idUser),
			IDProduct:        uint(idProduct),
			IDAlamat:         uint(idAlamat),
			IDKeranjang:      idKeranjang,
			MetodePengiriman: getString(p, "metode_pengiriman"),
			Pembayaran:       getString(p, "pembayaran"),
			Jumlah:           getInt(p, "jumlah"),
		}
		if err := db.DB.Create(&checkout).Error; err != nil {
			return nil, err
		}
		if err := db.DB.Preload("User").Preload("Product").Preload("Alamat").Preload("Keranjang").First(&checkout, checkout.IDCheckout).Error; err != nil {
			return nil, err
		}
		return checkout, nil
	},
}

var UpdateCheckout = &graphql.Field{
	Type: types.CheckoutType,
	Args: graphql.FieldConfigArgument{
		"id_checkout":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"id_product":        &graphql.ArgumentConfig{Type: graphql.Int},
		"id_alamat":         &graphql.ArgumentConfig{Type: graphql.Int},
		"id_keranjang":      &graphql.ArgumentConfig{Type: graphql.Int},
		"metode_pengiriman": &graphql.ArgumentConfig{Type: graphql.String},
		"pembayaran":        &graphql.ArgumentConfig{Type: graphql.String},
		"jumlah":            &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_checkout")
		var checkout models.Checkout
		if err := db.DB.First(&checkout, id).Error; err != nil {
			return nil, fmt.Errorf("checkout dengan id %d tidak ditemukan", id)
		}
		if v := getInt(p, "id_product"); v != 0 {
			checkout.IDProduct = uint(v)
		}
		if v := getInt(p, "id_alamat"); v != 0 {
			checkout.IDAlamat = uint(v)
		}
		if v := getInt(p, "id_keranjang"); v != 0 {
			u := uint(v)
			checkout.IDKeranjang = &u
		}
		if v := getString(p, "metode_pengiriman"); v != "" {
			checkout.MetodePengiriman = v
		}
		if v := getString(p, "pembayaran"); v != "" {
			checkout.Pembayaran = v
		}
		if v := getInt(p, "jumlah"); v != 0 {
			checkout.Jumlah = v
		}
		if err := db.DB.Save(&checkout).Error; err != nil {
			return nil, err
		}
		if err := db.DB.Preload("User").Preload("Product").Preload("Alamat").Preload("Keranjang").First(&checkout, checkout.IDCheckout).Error; err != nil {
			return nil, err
		}
		return checkout, nil
	},
}

var DeleteCheckout = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id_checkout": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_checkout")
		var checkout models.Checkout
		if err := db.DB.First(&checkout, id).Error; err != nil {
			return nil, fmt.Errorf("checkout dengan id %d tidak ditemukan", id)
		}
		if err := db.DB.Delete(&checkout).Error; err != nil {
			return nil, err
		}
		return "Checkout berhasil dihapus", nil
	},
}
