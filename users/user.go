package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	Username string
	Email    string
}

var counter int = 0
var usersList []UserResponse

func Register(router *gin.RouterGroup) {
	router.POST("/", Create)
	router.GET("/:username", Get)
	// router.PUT("/:id")
}

func Create(c *gin.Context) {

	var user UserResponse

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usersList = append(usersList, user)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"user":    user,
	})
	return
}

func Get(c *gin.Context) {

	username := c.Param("username")

	var res UserResponse
	for _, user := range usersList {
		if user.Username == username {
			res = user
			break
		}
	}

	c.JSON(http.StatusOK, res)
	return
}
