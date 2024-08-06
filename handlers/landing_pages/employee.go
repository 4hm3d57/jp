package landing_pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EmployeePage(c *gin.Context) {

	c.HTML(http.StatusOK, "employee.html", nil)

}
