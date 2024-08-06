package landing_pages

import (
	"job/db"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.JSON(http.StatusFound, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}

func EmployeeSession(c *gin.Context) {

	session := sessions.Default(c)
	userID := session.Get("userID").(string)
	oid, _ := primitive.ObjectIDFromHex(userID)

	user, err := db.GetUserID(oid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user data"})
		return
	}

	c.HTML(http.StatusOK, "employee.html", gin.H{"user": user})

}
