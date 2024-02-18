package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int64
	Username string
	Email    string
}

type UserList []User

var userList UserList

var counter int = 0

func Register(router *gin.RouterGroup) {
	router.POST("/", Create)
	router.GET("/:id", Get)
	router.PUT("/:id", Update)
	router.DELETE("/:id", Delete)
}

func Create(c *gin.Context) {

	var userInput struct {
		Username string
		Email    string
	}

	if err := c.Bind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	counter += 1

	var user User
	user.Id = int64(counter)
	user.Username = userInput.Username
	user.Email = userInput.Email

	userList.add(user)

	c.JSON(http.StatusCreated, user)
	return
}

func Get(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response := userList.get(id)

	c.JSON(http.StatusOK, response)
	return
}

func Update(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		fmt.Println("users/Update Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var userUpdateInput struct {
		Email string
	}

	c.Bind(&userUpdateInput)

	response := userList.update(id, userUpdateInput.Email)

	fmt.Println(userList)

	c.JSON(http.StatusOK, response)
	return
}

func Delete(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		fmt.Println("users/Delete Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userList.delete(id)

	c.JSON(http.StatusOK, "success")
	return
}
