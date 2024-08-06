package employee

import (
	"log"
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func ExperienceHandler(c *gin.Context) {

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

	institution := c.PostForm("institution")
	supervisor := c.PostForm("supervisor")
	telephone := c.PostForm("telephone")
	jobtitle := c.PostForm("jobtitle")
	start := c.PostForm("start")
	end := c.PostForm("end")
	duties := c.PostForm("duties")

	newExp := db.Experience{
		Institution: institution,
		Supervisor:  supervisor,
		Telephone:   telephone,
		Jobtitle:    jobtitle,
		Start:       start,
		End:         end,
		Duties:      duties,
	}

	err = db.InsertExperience(newExp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add experience data."})
		return
	}

	c.Redirect(http.StatusFound, "/experience")

}
