package landing_pages

import (
	"job/db"

	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Homepage(c *gin.Context) {
	session := sessions.Default(c)
	userIDstr, ok := session.Get("userID").(string)
	if !ok {
		c.JSON(500, gin.H{"error": "invalid user ID type"})
		log.Printf("User ID not found in session or invalid type")
		return
	}

	log.Printf("User ID from session: %s", userIDstr)

	userID, err := primitive.ObjectIDFromHex(userIDstr)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid user ID"})
		log.Printf("Error converting userID to ObjectID: %v", err)
		return
	}

	user, err := db.GetUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		log.Printf("Error fetching user from db: %v", err)
		return
	}

	profileURL := "/profile"
	switch user.AccountType {
	case "employer":
		profileURL = "/employer-profile"
	case "employee":
		profileURL = "/profile"
	}

	c.HTML(200, "home.html", gin.H{"profileURL": profileURL})
}
