package main

import (
	"log"
	"os"

	"github.com/PeterNex14/kioskecil-microservice/common/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Info: File .env tidak ditemukan, menggunakan env system")
	}

	dsn := database.DBSetupParams{
		DBHost: "localhost",
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: "users_service",
		DBPort: "5432",
	}

	_, err := database.InitDB(dsn)
	if err != nil {
		log.Fatal("Gagal koneksi ke database: ", err)
	}

}
