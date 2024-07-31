package handlers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func Homepage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", nil)

}