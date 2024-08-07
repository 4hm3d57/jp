package landing_pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingHomepage(c *gin.Context) {

	c.HTML(http.StatusOK, "landing_homepage.html", nil)

}
