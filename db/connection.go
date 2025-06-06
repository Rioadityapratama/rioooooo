package db

import (
	"fmt"
	"log"
	"os"

	//"bakulos_grapghql/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ File .env tidak ditemukan, lanjut pakai default")
	}

	// Ambil variabel dari file .env
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Format connection string (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	// Koneksi ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Gagal konek ke database: %v", err)
	}
	DB = database

	// AutoMigrate semua tabel
	//err = DB.AutoMigrate(
	//	&models.User{},
	//	&models.Penjual{},
	//	&models.Product{},
	//	&models.Keranjang{},
	//	&models.Checkout{},
	//	&models.Favorite{},
	//	&models.History{},
	//	&models.Search{},
	//)
	//if err != nil {
	//	log.Fatalf("❌ AutoMigrate gagal: %v", err)
	//}

	log.Println("✅ Koneksi database berhasil dan tabel dimigrasi!")
}
