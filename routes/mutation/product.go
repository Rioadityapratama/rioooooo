package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
)

var CreateProduct = &graphql.Field{
    Type: types.ProductType,
    Args: graphql.FieldConfigArgument{
        "id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        "name":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
        "kategori":   &graphql.ArgumentConfig{Type: graphql.String},
        "size":       &graphql.ArgumentConfig{Type: graphql.String},
        "deskripsi":  &graphql.ArgumentConfig{Type: graphql.String},
        "brand":      &graphql.ArgumentConfig{Type: graphql.String},
        "price":      &graphql.ArgumentConfig{Type: graphql.Int},
        "image":      &graphql.ArgumentConfig{Type: graphql.String},
        "warna":      &graphql.ArgumentConfig{Type: graphql.String},
        "stok":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        idPenjual := getInt(p, "id_penjual")
        var penjual models.Penjual
        if err := db.DB.First(&penjual, idPenjual).Error; err != nil {
            return nil, fmt.Errorf("penjual dengan id %d tidak ditemukan", idPenjual)
        }
        product := models.Product{
            IDPenjual: uint(idPenjual),
            Name:      getString(p, "name"),
            Kategori:  getString(p, "kategori"),
            Size:      getString(p, "size"),
            Deskripsi: getString(p, "deskripsi"),
            Brand:     getString(p, "brand"),
            Price:     p.Args["price"].(int),
            Image:     getString(p, "image"),
            Warna:     getString(p, "warna"),
            Stok:      p.Args["stok"].(int),
        }
        if err := db.DB.Create(&product).Error; err != nil {
            return nil, fmt.Errorf("gagal menyimpan produk: %v", err)
        }
        db.DB.Preload("Penjual").First(&product, product.IDProduct)
        return product, nil
    },
}

var UpdateProduct = &graphql.Field{
    Type: types.ProductType,
    Args: graphql.FieldConfigArgument{
        "id_product": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        "name":       &graphql.ArgumentConfig{Type: graphql.String},
        "size":       &graphql.ArgumentConfig{Type: graphql.String},
        "brand":      &graphql.ArgumentConfig{Type: graphql.String},
        "deskripsi":  &graphql.ArgumentConfig{Type: graphql.String},
        "kategori":   &graphql.ArgumentConfig{Type: graphql.String},
        "price":      &graphql.ArgumentConfig{Type: graphql.Int},
        "image":      &graphql.ArgumentConfig{Type: graphql.String},
        "warna":      &graphql.ArgumentConfig{Type: graphql.String},
        "stok":       &graphql.ArgumentConfig{Type: graphql.Int},
    },
    Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        id := getInt(p, "id_product")
        var product models.Product
        if err := db.DB.First(&product, id).Error; err != nil {
            return nil, fmt.Errorf("produk tidak ditemukan")
        }
        updates := map[string]interface{}{}
        if v := getString(p, "name"); v != "" {
            updates["name"] = v
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
        if v, ok := p.Args["price"].(int); ok {
            updates["price"] = int(v)
        }
        if v, ok := p.Args["stok"].(int); ok {
            updates["stok"] = int(v)
        }
        if err := db.DB.Model(&product).Updates(updates).Error; err != nil {
            return nil, fmt.Errorf("gagal update")
        }
        db.DB.Preload("Penjual").First(&product, id)
        return product, nil
    },
}

var DeleteProduct = &graphql.Field{
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
}
