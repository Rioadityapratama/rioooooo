package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var CreatePenjual = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"nama":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"telepon":  &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Args["password"].(string)), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		penjual := models.Penjual{
			Nama:     p.Args["nama"].(string),
			Email:    p.Args["email"].(string),
			Telepon:  getString(p, "telepon"),
			Password: string(hashedPassword),
		}
		if err := db.DB.Create(&penjual).Error; err != nil {
			return nil, err
		}
		return penjual, nil
	},
}

var UpdatePenjual = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"nama":       &graphql.ArgumentConfig{Type: graphql.String},
		"password":   &graphql.ArgumentConfig{Type: graphql.String},
		"email":      &graphql.ArgumentConfig{Type: graphql.String},
		"telepon":    &graphql.ArgumentConfig{Type: graphql.String},
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
		if v := getString(p, "telepon"); v != "" {
			updates["telepon"] = v
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
}

var UpdatePenjualProfil = &graphql.Field{
	Type: types.PenjualType,
	Args: graphql.FieldConfigArgument{
		"id_penjual": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"profil":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		IDPenjual := p.Args["id_penjual"].(int)
		profil := p.Args["profil"].(string)
		var penjual models.Penjual
		if err := db.DB.First(&penjual, IDPenjual).Error; err != nil {
			return nil, fmt.Errorf("penjual tidak ditemukan")
		}
		penjual.Profil = profil
		if err := db.DB.Save(&penjual).Error; err != nil {
			return nil, err
		}
		return penjual, nil
	},
}
