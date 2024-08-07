package employee

import (
	"log"
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func RefereesHandler(c *gin.Context) {

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

	name := c.PostForm("name")
	email := c.PostForm("email")
	title := c.PostForm("title")
	phone := c.PostForm("phone")
	institution := c.PostForm("institution")

	newReferee := db.Referee{
		Name:        name,
		Email:       email,
		Title:       title,
		Phone:       phone,
		Institution: institution,
	}

	err = db.InsertReferee(newReferee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding referee data"})
		return
	}

	c.Redirect(http.StatusFound, "/referees")
}
