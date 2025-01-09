package controller

import (
	"net/http"
	"strconv"
	"web/model"
	"web/validators"
	"github.com/gin-gonic/gin"
)

func UserSaveHandler (c *gin.Context) {
	var userInput model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	user := model.User {
		Username:userInput.Username,
		Email: userInput.Email,
		Password: userInput.Password,
	}
	
	jsonError, err := validators.ValidateUser(user)
	if (err != nil){
		c.JSON(400, jsonError)
		return
	}
	user.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "user has been created",
	})
}

func UserUpdateHandler (c *gin.Context) {
	var userInput model.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ :=  strconv.ParseInt(c.Param("id"), 10, 8)
	user := model.FindUserByID(id)
	
	if checkUser(user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	user.Id = id
	user.Username = userInput.Username
	user.Email = userInput.Email
	user.Password = userInput.Password
	
	errorsMap, err := validators.ValidateUser(user)
	if err != nil {
		c.JSON(400, errorsMap)
		return
	}
	user.Update()
	c.JSON(http.StatusOK, gin.H{
		"message": "user has been updated",
	})
}

func checkUser(user model.User) bool{
	return user == model.User{}
}

func UserDeleteHandler (c *gin.Context) {
	id, _ :=  strconv.ParseInt(c.Param("id"), 10, 8)
	user := model.FindUserByID(id)
	if checkUser(user) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}
	user.Id = id
	user.Delete()
	c.JSON(http.StatusOK, gin.H{
		"message": "user has been deleted successfully",
	})
}