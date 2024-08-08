package main

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	employeeHandler "job/handlers/employee"
	employerHandler "job/handlers/employer"
	landingPages "job/handlers/landing_pages"
	register "job/handlers/register"
)

func main() {

	r := gin.Default()

	// create new session store using cookies
	store := cookie.NewStore([]byte("secret_key"))
	// applies middleware to router 'r'
	r.Use(sessions.Sessions("session_id", store))

	r.Static("/static", "./templates")

	r.LoadHTMLGlob("templates/*.html")

	r.Static("/register", "./templates/register")

	r.GET("/", landingPages.LandingHomepage)

	r.GET("/login", func(c *gin.Context) {
		c.File("templates/register/login.html")
	})

	r.POST("/login", register.LoginHandler)

	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/register/signup.html")
	})

	r.POST("/signup", register.SignupHandler)

	r.GET("/home", landingPages.Homepage)
	r.GET("/jobs", landingPages.JobPage)
	r.GET("/employer", landingPages.EmployerPage)
	r.GET("/employee", landingPages.AuthMiddleWare(), landingPages.EmployeeSession)

	r.GET("/contact", func(c *gin.Context) {
		c.File("templates/contact.html")
	})

	r.POST("/contact", landingPages.ContactPage)

	// profile
	r.GET("/profile", func(c *gin.Context) {
		c.File("templates/employee/profile.html")
	})

	r.POST("/profile", employeeHandler.AuthMiddleWare(), employeeHandler.ProfileHandler, employeeHandler.ProfileSession)

	// academics
	r.GET("/academics", func(c *gin.Context) {
		c.File("templates/employee/academics.html")
	})

	r.POST("/academics", employeeHandler.AcademicsHandler)

	// experience
	r.GET("/experience", func(c *gin.Context) {
		c.File("templates/employee/experience.html")
	})

	r.POST("/experience", employeeHandler.ExperienceHandler)

	// language
	r.GET("/language", func(c *gin.Context) {
		c.File("templates/employee/language.html")
	})

	r.POST("/language", employeeHandler.LanguageHandler)

	// profession
	r.GET("/profession", func(c *gin.Context) {
		c.File("templates/employee/profession.html")
	})

	r.POST("/profession", employeeHandler.ProfessionHandler)

	// referees
	r.GET("/referees", func(c *gin.Context) {
		c.File("templates/employee/referees.html")
	})

	r.POST("/referee", employeeHandler.RefereesHandler)

	//training
	r.GET("/training", func(c *gin.Context) {
		c.File("templates/employee/training_workshop.html")
	})

	r.POST("/training", employeeHandler.TrainingHandler)

	// employer profile
	r.GET("/employer-profile", func(c *gin.Context) {
		c.File("templates/employer/profile.html")
	})

	r.POST("/employer-profile", employerHandler.EmployerProfileHandler)

	// post job
	r.GET("/post-job", func(c *gin.Context) {
		c.File("templates/employer/postjob.html")
	})

	r.POST("/post-job", employerHandler.PostJobHandler)

	// logout
	r.GET("/logout", employeeHandler.LogoutHandler)

	log.Fatal(r.Run(":9000"))
}
