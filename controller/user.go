package controller

import (
	"chefling/app"
	"chefling/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := app.User{}
	c.ShouldBind(&user)
	token, ok := user.Create()
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": token,
		})
		return
	}
	c.JSON(200, gin.H{
		"success":   true,
		"authToken": token,
	})
}

func LoginUser(c *gin.Context) {
	user := app.User{}
	c.ShouldBind(&user)
	token, ok := app.Login(user.Email, user.Password)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": token,
		})
		return
	}
	c.JSON(200, gin.H{
		"success":   true,
		"authToken": token,
	})
}

func GetUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	authToken := strings.Split(authHeader, " ")[1]
	claims, ok := utils.ExtractClaims(authToken)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	}
	email := claims["Email"]
	user, ok := app.GetUser(email)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": "User doesn't exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"user":    user,
	})

}

func EditUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	authToken := strings.Split(authHeader, " ")[1]
	claims, ok := utils.ExtractClaims(authToken)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	}
	email := claims["Email"]
	updateUser := &app.User{}
	c.ShouldBind(&updateUser)
	message, ok := updateUser.EditUser(email)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"message": "User doesn't exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"user":    message,
	})

}
