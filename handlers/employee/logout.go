package employee

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear() // clear all session data
	session.Save()  // save changes to the session store

	c.Redirect(http.StatusSeeOther, "/")
}

func CacheControlMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		c.Header("Cache-Control", "post-check=0, pre-check=0")
		c.Header("Pragma", "no-cache")
		c.Next()
	}
}
