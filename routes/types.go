package routes

import (
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id_user": &graphql.Field{Type: graphql.Int},
		"nama":    &graphql.Field{Type: graphql.String},
		"email":   &graphql.Field{Type: graphql.String},
		"no_telp": &graphql.Field{Type: graphql.String},
	},
})

var CheckoutType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Checkout",
	Fields: graphql.Fields{
		"id_checkout": &graphql.Field{
			Type: graphql.Int,
		},
		"id_user": &graphql.Field{
			Type: graphql.Int,
		},
		"id_product": &graphql.Field{
			Type: graphql.Int,
		},
		"id_keranjang": &graphql.Field{
			Type: graphql.Int,
		},
		"id_alamat": &graphql.Field{
			Type: graphql.Int,
		},
		"metode_pengiriman": &graphql.Field{
			Type: graphql.String,
		},
		"pembayaran": &graphql.Field{
			Type: graphql.String,
		},
		"jumlah": &graphql.Field{
			Type: graphql.Int,
		},

		// RELASI
		"user": &graphql.Field{
			Type: UserType,
		},
		"product": &graphql.Field{
			Type: ProductType,
		},
		"keranjang": &graphql.Field{
			Type: KeranjangType,
		},
		"alamat": &graphql.Field{
			Type: AlamatType,
		},
	},
})

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id_product": &graphql.Field{Type: graphql.Int},
		"id_penjual": &graphql.Field{Type: graphql.Int},
		"namaproduk": &graphql.Field{Type: graphql.String},
		"kategori":   &graphql.Field{Type: graphql.String},
		"size":       &graphql.Field{Type: graphql.String},
		"deskripsi":  &graphql.Field{Type: graphql.String},
		"brand":      &graphql.Field{Type: graphql.String},
		"price":      &graphql.Field{Type: graphql.Float},
		"image":      &graphql.Field{Type: graphql.String},
		"warna":      &graphql.Field{Type: graphql.String},
		"penjual":    &graphql.Field{Type: PenjualType},
	},
})

var KeranjangType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Keranjang",
	Fields: graphql.Fields{
		"id_keranjang": &graphql.Field{Type: graphql.Int},
		"id_user":      &graphql.Field{Type: graphql.Int},
		"id_product":   &graphql.Field{Type: graphql.Int},
		"jumlah":       &graphql.Field{Type: graphql.Int},
		"user": &graphql.Field{
			Type: UserType,
		},
		"product": &graphql.Field{
			Type: ProductType,
		 },

	},
})

var FavoriteType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Favorite",
	Fields: graphql.Fields{
		"id_favorite": &graphql.Field{Type: graphql.Int},
		"id_product":  &graphql.Field{Type: graphql.Int},
		"id_user":     &graphql.Field{Type: graphql.Int},
		"user": &graphql.Field{
			Type: UserType,
		},
		"product": &graphql.Field{
			Type: ProductType,
		},
	},
})

var HistoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "History",
	Fields: graphql.Fields{
		"id_history":  &graphql.Field{Type: graphql.Int},
		"id_checkout": &graphql.Field{Type: graphql.Int},
		"checkout": &graphql.Field{
			Type: CheckoutType,
		},
	},
})


var PenjualType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Penjual",
	Fields: graphql.Fields{
		"id_penjual": &graphql.Field{Type: graphql.Int},
		"nama":       &graphql.Field{Type: graphql.String},
		"email":      &graphql.Field{Type: graphql.String},
	},
})

var AlamatType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Alamat",
	Fields: graphql.Fields{
		"id_alamat": &graphql.Field{Type: graphql.Int},
		"id_user":   &graphql.Field{Type: graphql.Int},
		"alamat":    &graphql.Field{Type: graphql.String},
		"user": &graphql.Field{
			Type: UserType,
		},
	},
})
