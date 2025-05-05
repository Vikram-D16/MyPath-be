package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Vikram-D16/MyPath-be/models"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	var err error

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Connected to PostgreSQL with GORM")
			break
		}
		log.Printf("❌ DB connection failed (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to database after retries: %v", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Auto-migrate failed: %v", err)
	}
}
