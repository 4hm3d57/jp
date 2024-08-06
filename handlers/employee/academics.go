package employee

import (
	"job/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AcademicsHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("error parsing form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		return
	}

	ed_level := c.PostForm("ed_level")
	institution := c.PostForm("institution")
	course := c.PostForm("course")
	timeframe := c.PostForm("timeframe")

	log.Printf("Added data => ed_level: %s, institution: %s, course: %s, timeframe: %s", ed_level, institution, course, timeframe)

	if ed_level == "" || institution == "" || course == "" || timeframe == "" {
		log.Printf("error adding data: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newAcademics := db.Academics{
		Ed_level:    ed_level,
		Institution: institution,
		Course:      course,
		Timeframe:   timeframe,
	}

	err = db.InsertAcademics(newAcademics)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding academics data."})
		return
	}

	c.Redirect(http.StatusFound, "/")

}
