package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ File .env tidak ditemukan, lanjut pakai default")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke database: %v", err)
	}

	// Aktifkan logging query SQL (opsional untuk debug)
	DB = database.Debug()

	log.Println("✅ Koneksi database berhasil!")
}
