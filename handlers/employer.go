package handlers


import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func EmployerPage(c *gin.Context) {

	c.HTML(http.StatusOK, "employer.html", nil);

}