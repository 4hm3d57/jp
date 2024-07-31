package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func JobPage(c *gin.Context) {

	c.HTML(http.StatusOK, "jobs.html", nil)
	
}