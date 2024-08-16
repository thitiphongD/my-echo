package database

import (
	"fmt"
	"log"
	"os/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}
	fmt.Println("Database connection established")
	DB = db
	err = db.AutoMigrate(&user.User{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}
	fmt.Println("Database migration completed")
	return db, nil
}
