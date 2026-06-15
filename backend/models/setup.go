package models

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

// Tambahkan parameter fresh (true/false) di fungsi ini
func ConnectDatabase(fresh bool) {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: no .env file found, using system environment variable")
    }

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    if fresh {
        log.Println("Mengosongkan database (migrate:fresh)...")
        err := database.Migrator().DropTable(&Book{}, &User{})
        if err != nil {
            log.Printf("Gagal drop tabel: %v", err)
        } else {
            log.Println("Semua tabel berhasil di-drop.")
        }
    }

    log.Println("⚙ Running database migration...")
    err = database.AutoMigrate(&Book{}, &User{})
    if err != nil {
        log.Fatalf("Gagal auto migrate: %v", err)
    }

    DB = database
    log.Println("Database connected successfully")
}