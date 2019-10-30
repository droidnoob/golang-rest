package main

import (
	"chefling/app"
	"chefling/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app.DBInit()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/user/signup", controller.CreateUser)
	r.POST("/user/signin", controller.LoginUser)
	r.Use(app.Auth(os.Getenv("jwt_secret_password")))
	r.GET("/user/profile", controller.GetUser)
	r.PATCH("/user/profile/update", controller.EditUser)
	fmt.Println("server running")
	r.Run()
}
