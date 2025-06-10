package mutation

import (
	"bakulos_grapghql/db"
	"bakulos_grapghql/models"
	"bakulos_grapghql/routes/types"
	"fmt"

	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

var CreateUser = &graphql.Field{
	Type: types.UserType,
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
		user := models.User{
			Nama:     p.Args["nama"].(string),
			Email:    p.Args["email"].(string),
			Telepon:  getString(p, "telepon"),
			Password: string(hashedPassword),
		}
		if err := db.DB.Create(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	},
}

var UpdateUser = &graphql.Field{
	Type: types.UserType,
	Args: graphql.FieldConfigArgument{
		"id_user":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"nama":     &graphql.ArgumentConfig{Type: graphql.String},
		"email":    &graphql.ArgumentConfig{Type: graphql.String},
		"telepon":  &graphql.ArgumentConfig{Type: graphql.String},
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
		if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
			return nil, fmt.Errorf("gagal update data user: %v", err)
		}
		db.DB.First(&user, id)
		return user, nil
	},
}
