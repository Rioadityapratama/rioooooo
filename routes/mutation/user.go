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
		email := p.Args["email"].(string)

		// ✅ Cek apakah email sudah ada di database
		var existingUser models.User
		if err := db.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
			return nil, fmt.Errorf("email sudah terdaftar")
		}

		// ✅ Kalau belum ada, baru hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Args["password"].(string)), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		user := models.User{
			Nama:     p.Args["nama"].(string),
			Email:    email,
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
		"id_user":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"nama":         &graphql.ArgumentConfig{Type: graphql.String},
		"telepon":      &graphql.ArgumentConfig{Type: graphql.String},
		"password":     &graphql.ArgumentConfig{Type: graphql.String},      // password baru
		"old_password": &graphql.ArgumentConfig{Type: graphql.String},      // password lama
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := getInt(p, "id_user")
		var user models.User

		if err := db.DB.First(&user, id).Error; err != nil {
			return nil, fmt.Errorf("user tidak ditemukan")
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
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
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

		if err := db.DB.Model(&user).Updates(updates).Error; err != nil {
			return nil, err
		}

		db.DB.First(&user, id)
		return user, nil
	},
}

var UpdateUserProfil = &graphql.Field{
	Type: types.UserType,
	Args: graphql.FieldConfigArgument{
		"id_user": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		"profil":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		IDUser := p.Args["id_user"].(int)
		profil := p.Args["profil"].(string)
		var user models.User
		if err := db.DB.First(&user, IDUser).Error; err != nil {
			return nil, fmt.Errorf("user tidak ditemukan")
		}
		user.Profil = profil
		if err := db.DB.Save(&user).Error; err != nil {
			return nil, err
		}
		return user, nil
	},
}

var LoginUser = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "LoginUserResponse",
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

		var user models.User
		if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return nil, fmt.Errorf("email tidak ditemukan")
		}

		// ✅ Inilah kunci validasi yang benar
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return nil, fmt.Errorf("password salah")
		}

		return map[string]interface{}{
			"message": "Login berhasil",
		}, nil
	},
}
//ini user yang belum masuk

var ForgetPasswordUser = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"email":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"new_password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		email := getString(p, "email")
		newPassword := getString(p, "new_password")

		// ✅ Cek apakah email ada
		var user models.User
		if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
			return nil, fmt.Errorf("email tidak ditemukan")
		}

		// ✅ Hash password baru
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("gagal hash password baru: %v", err)
		}

		// ✅ Update password langsung
		if err := db.DB.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
			return nil, fmt.Errorf("gagal update password: %v", err)
		}

		return "Password berhasil diganti", nil
	},
}


















































