package main

import (
	"job/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/static", "./templates/css")

	r.LoadHTMLGlob("templates/*.html")

	r.Static("/register", "./templates/register")

	r.GET("/", func(c *gin.Context) {
		c.File("templates/register/login.html")
	})

	r.POST("/login", handlers.LoginHandler)

	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/register/signup.html")
	})

	r.POST("/signup", handlers.SignupHandler)



	log.Fatal(r.Run(":8080"))
}
