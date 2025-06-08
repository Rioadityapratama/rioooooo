package db

import (
	"bakulos_grapghql/models"
	"log"
)

func AutoMigrateTables() {
	migrations := []struct {
		Name  string
		Model interface{}
	}{
		{"User", &models.User{}},
		{"User_Penjual", &models.Penjual{}},
		{"Alamat", &models.Alamat{}},
		{"Product", &models.Product{}},
		{"Keranjang", &models.Keranjang{}},
		{"Checkout", &models.Checkout{}},
		{"Favorite", &models.Favorite{}},
		{"History", &models.History{}},
	}

	for _, m := range migrations {
		log.Printf("🚧 Migrasi tabel %s...", m.Name)
		if err := DB.AutoMigrate(m.Model); err != nil {
			log.Fatalf("❌ AutoMigrate %s gagal: %v", m.Name, err)
		}
		log.Printf("✅ Migrasi tabel %s berhasil", m.Name)
	}

	log.Println("🎉 Semua tabel berhasil dimigrasi!")
}
