package common

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=SocialMediaApp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB connect error: ", err)
	}
	fmt.Println("DB connected! ", db)

	DB = db

	return DB
}

func GetDB() *gorm.DB {
	return DB
}
