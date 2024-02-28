package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"nishin/social-media-API/users"

	"nishin/social-media-API/common"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	common.Init() // Initialize DB connection

	users.AutoMigrate()

	r := gin.Default()

	v1 := r.Group("/api/v1")

	users.Register(v1.Group("/users"))

	r.Run("localhost:5000")
}
