package landing_pages

import (
	"job/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContactPage(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "bad method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("Error parsing form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	log.Printf("Data received => name: %s, email: %s, message: %s", name, email, message)

	if name == "" || email == "" || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	newContact := db.Contacts{
		Name:    name,
		Email:   email,
		Message: message,
	}

	err = db.InsertContact(newContact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding contacts"})
	}

	c.HTML(http.StatusOK, "contact.html", nil)

}
