package handlers

import (
	"job/db"
	"log"
	"net/http"

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
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	acc_type := c.PostForm("acc_type")

	log.Printf("Added data => name: %s, email: %s, password: %s, acc_type: %s", name, email, password, acc_type)

	if name == "" || email == "" || password == "" || acc_type == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "all fields are required"})
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

	// TODO: redirect to the assigned page

	switch newUser.AccountType {
	case "employer":
		c.Redirect(http.StatusSeeOther, "/home")
	case "employee":
		c.Redirect(http.StatusSeeOther, "/home")
	}

}
