package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"
	"github.com/graphql-go/graphql"
)

var CreateHistory = &graphql.Field{
	Type: types.HistoryType,
	Args: graphql.FieldConfigArgument{
		"id_checkout": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idCheckout := getInt(p, "id_checkout")
		var checkout models.Checkout
		if err := db.DB.First(&checkout, idCheckout).Error; err != nil {
			return nil, fmt.Errorf("checkout dengan id %d tidak ditemukan", idCheckout)
		}
		history := models.History{
			IDCheckout: uint(idCheckout),
		}
		if err := db.DB.Create(&history).Error; err != nil {
			return nil, fmt.Errorf("gagal membuat history: %v", err)
		}
		if err := db.DB.Preload("Checkout").Preload("Checkout.User").Preload("Checkout.Product").Preload("Checkout.Alamat").Preload("Checkout.Keranjang").First(&history, history.IDHistory).Error; err != nil {
			return nil, err
		}
		return history, nil
	},
}

var DeleteHistory = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id_history": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_history")
		var history models.History
		if err := db.DB.First(&history, id).Error; err != nil {
			return nil, fmt.Errorf("history dengan id %d tidak ditemukan", id)
		}
		if err := db.DB.Delete(&history).Error; err != nil {
			return nil, fmt.Errorf("gagal menghapus history: %v", err)
		}
		return "History berhasil dihapus", nil
	},
}
