package users

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup) {
	router.GET("/:id")
	router.POST("/")
	router.PUT("/:id")
}
