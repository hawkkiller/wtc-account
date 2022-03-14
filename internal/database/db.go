package database

import (
	"fmt"
	"github.com/hawkkiller/wtc-account/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetupDB() {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PASSWORD"))))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	DB = db

	err = DB.AutoMigrate(&model.UserProfile{})

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
