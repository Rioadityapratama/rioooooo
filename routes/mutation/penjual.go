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
		email := p.Args["email"].(string)

		// ✅ Cek apakah email sudah ada di database
		var existingPenjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&existingPenjual).Error; err == nil {
			return nil, fmt.Errorf("email sudah terdaftar")
		}

		// ✅ Kalau belum ada, baru hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Args["password"].(string)), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		penjual := models.Penjual{
			Nama:     p.Args["nama"].(string),
			Email:    email,
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
		"id_penjual":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"nama":         &graphql.ArgumentConfig{Type: graphql.String},
		"telepon":      &graphql.ArgumentConfig{Type: graphql.String},
		"password":     &graphql.ArgumentConfig{Type: graphql.String},    
		"old_password": &graphql.ArgumentConfig{Type: graphql.String},    
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_penjual")
		var penjual models.Penjual

		if err := db.DB.First(&penjual, id).Error; err != nil {
			return nil, fmt.Errorf("penjual tidak ditemukan")
		}

		updates := map[string]interface{}{}

		// Update nama & telepon tetap seperti biasa
		if v := getString(p, "nama"); v != "" {
			updates["nama"] = v
		}
		if v := getString(p, "telepon"); v != "" {
			updates["telepon"] = v
		}

		// ✅ Change Password Section
		if v := getString(p, "password"); v != "" {
			// Password baru dimasukkan → maka harus validasi old_password dulu
			oldPassword := getString(p, "old_password")
			if oldPassword == "" {
				return nil, fmt.Errorf("password lama wajib diisi")
			}

			// Validasi apakah old_password cocok
			err := bcrypt.CompareHashAndPassword([]byte(penjual.Password), []byte(oldPassword))
			if err != nil {
				return nil, fmt.Errorf("password lama salah")
			}

			// Hash password baru
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.DefaultCost)
			if err != nil {
				return nil, err
			}
			updates["password"] = string(hashedPassword)
		}

		if err := db.DB.Model(&penjual).Updates(updates).Error; err != nil {
			return nil, err
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

var LoginPenjual = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "LoginPenjualResponse",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
		},
	}),
	Args: graphql.FieldConfigArgument{
		"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := getString(p, "email")
		password := getString(p, "password")

		var penjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&penjual).Error; err != nil {
			return nil, fmt.Errorf("email tidak ditemukan")
		}

		// ✅ Inilah kunci validasi yang benar
		err := bcrypt.CompareHashAndPassword([]byte(penjual.Password), []byte(password))
		if err != nil {
			return nil, fmt.Errorf("password salah")
		}

		return map[string]interface{}{
			"message": "Login berhasil",
		}, nil
	},
}

var ForgetPasswordPenjual = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"email":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"new_password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := getString(p, "email")
		newPassword := getString(p, "new_password")

		// ✅ Cek apakah email ada
		var penjual models.Penjual
		if err := db.DB.Where("email = ?", email).First(&penjual).Error; err != nil {
			return nil, fmt.Errorf("email tidak ditemukan")
		}

		// ✅ Hash password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("gagal hash password baru: %v", err)
		}

		// ✅ Update password langsung
		if err := db.DB.Model(&penjual).Update("password", string(hashedPassword)).Error; err != nil {
			return nil, fmt.Errorf("gagal update password: %v", err)
		}

		return "Password berhasil diganti", nil
	},
}
