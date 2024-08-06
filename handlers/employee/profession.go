package employee

import (
	"log"
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func ProfessionHandler(c *gin.Context) {

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

	county := c.PostForm("county")
	institution := c.PostForm("institution")
	course := c.PostForm("course")
	timeframe := c.PostForm("timeframe")

	newProffesion := db.Profession{
		County:      county,
		Institution: institution,
		Course:      course,
		Timeframe:   timeframe,
	}

	err = db.InsertProfession(newProffesion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding profession data"})
		return
	}

	c.Redirect(http.StatusFound, "/profession")
}
