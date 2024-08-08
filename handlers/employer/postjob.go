package employer

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"job/db"
)

func PostJobHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("error parsing form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data."})
		return
	}

	title := c.PostForm("title")
	city := c.PostForm("city")
	county := c.PostForm("county")
	category := c.PostForm("category")
	date := c.PostForm("date")
	_type := c.PostForm("_type")
	experience := c.PostForm("experience")
	description := c.PostForm("description")
	responsibility := c.PostForm("responsibility")
	requirements := c.PostForm("requirements")


	postjob := db.Postjob {
		Title: title,
		City: city,
		County: county,
		Category: category,
		Date: date,
		Type: _type,
		Experience: experience,
		Description: description,
		Responsibility: responsibility,
		Requirements: requirements,
	}


	err = db.InsertJob(postjob)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "error adding job data"})
		return
	}

	c.Redirect(http.StatusFound, "/postjob")

}
