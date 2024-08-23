package register

import (
	"job/db"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		log.Printf("error parsing form data: %v", err)
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	acc_type := c.PostForm("acc_type")

	log.Printf("Received data => name: %s, email: %s, password: %s, acc_type: %s", name, email, password, acc_type)

	if name == "" || email == "" || password == "" || acc_type == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "all fields are required"})
		return
	}

	newUser := db.User{
		Name:        name,
		Email:       email,
		Password:    password,
		AccountType: acc_type,
	}

	err = db.InsertUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding user"})
		log.Printf("error adding user to db: %v", err)
		return
	}

	// Verify that the ID is properly set
	log.Printf("New user ID: %s", newUser.ID.Hex())

	session := sessions.Default(c)
	session.Set("userID", newUser.ID.Hex())
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		log.Printf("error saving session: %v", err)
		return
	}

	log.Printf("Session data after signup: %s", session.Get("userID"))

	// Redirect after signup
	switch newUser.AccountType {
	case "employer":
		c.Redirect(http.StatusSeeOther, "/employer-profile")
	case "employee":
		c.Redirect(http.StatusSeeOther, "/profile")
	}
}
