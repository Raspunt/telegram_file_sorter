package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func InitDB() *gorm.DB{

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db_str := os.Getenv("DB_connection_str")
	fmt.Println(db_str)

	db, err := gorm.Open(postgres.Open(db_str), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Lesson{})
	db.AutoMigrate(&File{})

	return db

}
