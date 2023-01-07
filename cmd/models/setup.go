package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"))

	if err != nil {
		panic("Failed connect to database")
	}

	database.AutoMigrate(&Role{})
	database.AutoMigrate(&Permission{})
	database.AutoMigrate(&Book{})
	database.AutoMigrate(&User{})

	DB = database
}
