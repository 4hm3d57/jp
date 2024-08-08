package employer

import (
	"job/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmployerProfileHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("error parsing form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		return
	}

	name := c.PostForm("name")
	est := c.PostForm("est")
	_type := c.PostForm("_type")
	people := c.PostForm("people")
	web := c.PostForm("web")
	city := c.PostForm("city")
	street := c.PostForm("street")
	county := c.PostForm("county")
	zip := c.PostForm("zip")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	background := c.PostForm("background")
	service := c.PostForm("service")
	expertise := c.PostForm("experience")

	emp_profile := db.EmployerProfile{
		Name:        name,
		Established: est,
		Type:        _type,
		People:      people,
		Website:     web,
		City:        city,
		Street:      street,
		County:      county,
		Zip:         zip,
		Phone:       phone,
		Email:       email,
		Background:  background,
		Service:     service,
		Expertise:   expertise,
	}

	err = db.InsertEmployerProfile(emp_profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding profile"})
		return
	}

	c.Redirect(http.StatusFound, "/employer-profile")

}
