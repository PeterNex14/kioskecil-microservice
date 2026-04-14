package main

import (
	"log"

	"github.com/PeterNex14/kioskecil-microservice/common/config"
	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/joho/godotenv"
)

func main() {
	// Mencoba load .env dari beberapa lokasi yang mungkin
	_ = godotenv.Load(".env")    // Mencoba di folder saat ini
	_ = godotenv.Load("../.env") // Mencoba di folder parent (untuk lokal)

	db_host, err := config.GetEnv("DB_HOST")
	if err != nil {
		log.Fatal("DBHost tidak ditemukan")
	}

	db_user, err := config.GetEnv("DB_USER")
	if err != nil {
		log.Fatal("DBHost tidak ditemukan")
	}

	db_password, err := config.GetEnv("DB_PASSWORD")
	if err != nil {
		log.Fatal("DBHost tidak ditemukan")
	}

	db_name, err := config.GetEnv("DB_NAME")
	if err != nil {
		log.Fatal("DBHost tidak ditemukan")
	}

	db_port, err := config.GetEnv("DB_PORT")
	if err != nil {
		log.Fatal("DBHost tidak ditemukan")
	}

	dsn := database.DBSetupParams{
		DBHost:     db_host,
		DBUser:     db_user,
		DBPassword: db_password,
		DBName:     db_name,
		DBPort:     db_port,
	}

	_, err = database.InitDB(dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

	log.Println("Berhasil terhubung ke database")
}
