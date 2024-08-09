package landing_pages

import (
	"job/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Homepage(c *gin.Context) {

	session := sessions.Default(c)
	userIDstr, ok := session.Get("userID").(string)
	if !ok {
		c.JSON(500, gin.H{"error": "invalid user ID type"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDstr)
	if err != nil {
		c.JSON(500, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := db.GetUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	profileURL := "/profile"
	switch user.AccountType {
	case "employer":
		profileURL = "/employer-profile"
	case "employee":
		profileURL = "/profile"

	}

	// render the home page
	c.HTML(200, "home.html", gin.H{"profileURL": profileURL})

}
