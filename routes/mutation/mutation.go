package mutation

import (
	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// === USER ===
		"createUser": CreateUser,
		"updateUser": UpdateUser,
		// === PENJUAL ===
		"createPenjual":       CreatePenjual,
		"updatePenjual":       UpdatePenjual,
		"updatePenjualProfil": UpdatePenjualProfil,
		// === PRODUCT ===
		"createProduct": CreateProduct,
		"updateProduct": UpdateProduct,
		"deleteProduct": DeleteProduct,
		// === KERANJANG ===
		"createKeranjang": CreateKeranjang,
		"updateKeranjang": UpdateKeranjang,
		"deleteKeranjang": DeleteKeranjang,
		// === FAVORITE ===
		"createFavorite": CreateFavorite,
		"updateFavorite": UpdateFavorite,
		"deleteFavorite": DeleteFavorite,
		// === ALAMAT ===
		"createAlamat": CreateAlamat,
		"updateAlamat": UpdateAlamat,
		"deleteAlamat": DeleteAlamat,
		// === CHECKOUT ===
		"createCheckout": CreateCheckout,
		"updateCheckout": UpdateCheckout,
		"deleteCheckout": DeleteCheckout,
		// === HISTORY ===
		"createHistory": CreateHistory,
		"deleteHistory": DeleteHistory,
	},
})
