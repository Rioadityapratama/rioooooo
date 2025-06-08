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
		{"Product", &models.Product{}},
		{"Keranjang", &models.Keranjang{}},
		{"Search", &models.Search{}},
		{"Checkout", &models.Checkout{}},
		{"History", &models.History{}},
		{"Favorite", &models.Favorite{}},
	}

	for _, m := range migrations {
		if err := DB.AutoMigrate(m.Model); err != nil {
			log.Fatalf("❌ AutoMigrate %s gagal: %v", m.Name, err)
		}
		log.Printf("✅ Migrasi tabel %s berhasil", m.Name)
	}

	log.Println("🎉 Semua tabel berhasil dimigrasi!")
}
