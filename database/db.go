package database

import (
	"github.com/NochboolPrime/graphql-blog/config"
	"github.com/NochboolPrime/graphql-blog/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("postgres", config.GetDBURL())
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&models.Post{}, &models.Comment{})
}
