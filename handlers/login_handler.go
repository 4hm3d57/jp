package handlers

import (
	"job/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "bad request method"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		log.Printf("error parsing form data: %v", err)
		return
	}

	email := c.PostForm("email")
	password := c.PostForm("password")

	log.Printf("Received data => email: %s, password: %s", email, password)

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields are required"})
		return
	}

	user, err := db.GetUser(email, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error retrieving user"})
		log.Printf("Error retrieving user: %v", err)
		return
	}

	// shit code will remove
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// TODO: redirect to the assigned page
	switch user.AccountType {
	case "employer":
		c.Redirect(http.StatusSeeOther, "/home")
	case "employee":
		c.Redirect(http.StatusSeeOther, "/home")
	}
}
