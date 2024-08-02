package main

import (
	"job/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/static", "./templates")

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

	r.GET("/home", handlers.Homepage)
	r.GET("/jobs", handlers.JobPage)
	r.GET("/employer", handlers.EmployerPage)
	r.GET("/employee", handlers.EmployeePage)

	r.GET("/contact", func(c *gin.Context) {
		c.File("templates/contact.html")
	})

	r.POST("/contact", handlers.ContactPage)

	r.GET("/profile", func(c *gin.Context) {
		c.File("templates/employee/profile.html")
	})
	r.POST("/profile", handlers.ProfileHandler)

	log.Fatal(r.Run(":9000"))
}
