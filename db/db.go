package db

import (
	"fmt"
	"os"

	"ask-anon-ques/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB(){
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to db: %v\n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Database started...")

	// DB.Migrator().DropTable(&models.Question{},&models.User{})

	err := DB.AutoMigrate(&models.User{},&models.Question{})
	if err != nil {
		panic("Failed to migrate")
	}

	fmt.Println("Migration success.")
}