package landing_pages

import (
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func EmployeePage(c *gin.Context) {

	employee, err := db.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting users"})
		return
	}

	c.HTML(http.StatusOK, "employee.html", gin.H{"employee": employee})

}
