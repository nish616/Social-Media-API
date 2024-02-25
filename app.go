package main

import (
	"github.com/gin-gonic/gin"

	"nishin/social-media-API/users"

	"nishin/social-media-API/common"
)

func main() {

	common.Init() // Initialize DB connection

	r := gin.Default()

	v1 := r.Group("/api/v1")

	users.Register(v1.Group("/users"))
	users.AutoMigrate()

	r.Run("localhost:5000")
}
