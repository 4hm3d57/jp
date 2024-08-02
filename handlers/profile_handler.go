package handlers

import (
	"github.com/gin-gonic/gin"
	"job/db"
	"log"
	"net/http"
)

func ProfileHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("failed to parse form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse form data"})
		return
	}

	f_name := c.PostForm("f_name")
	l_name := c.PostForm("l_name")
	born := c.PostForm("born")
	email := c.PostForm("email")
	ed_level := c.PostForm("ed_level")
	ed_course := c.PostForm("ed_course")
	gender := c.PostForm("gender")
	city := c.PostForm("city")
	street := c.PostForm("street")
	zip := c.PostForm("zip")
	county := c.PostForm("county")
	phone := c.PostForm("phone")
	about := c.PostForm("about")

	log.Printf("first_name: %s, last_name: %s, born: %s, email: %s, ed_level: %s, ed_course: %s, gender: %s, city: %s, street: %s, zip: %s, county: %s, phone: %s",
		f_name, l_name, born, email, ed_level, ed_course, gender, city, street, zip, county, phone)

	if f_name == "" || l_name == "" || born == "" || email == "" || ed_level == "" || ed_course == "" || gender == "" || city == "" || street == "" || zip == "" || county == "" || phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newProfile := db.Profile{
		First_name:   f_name,
		Last_name:    l_name,
		Born:         born,
		Email:        email,
		Ed_Level:     ed_level,
		Ed_course:    ed_course,
		Gender:       gender,
		City:         city,
		Street:       street,
		Zip:          zip,
		County:       county,
		Phone_number: phone,
		About:        about,
	}

	err = db.InsertProfile(newProfile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding user profile"})
		log.Printf("error adding user profile: %s", err)
		return
	}

}
