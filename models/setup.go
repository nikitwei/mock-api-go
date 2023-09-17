package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/mock_api_go"))
	if err != nil {
		fmt.Println("Failed Connect to DB")
	}

	db.AutoMigrate(&User{})
	DB = db
}
