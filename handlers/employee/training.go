package employee

import (
	"log"
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func TrainingHandler(c *gin.Context) {

	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid method request"})
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Printf("error parsing form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error parsing form data"})
		return
	}

	training := c.PostForm("training")
	institution := c.PostForm("institution")
	timeframe := c.PostForm("timeframe")

	newTraining := db.Train{
		Training:    training,
		Institution: institution,
		Timeframe:   timeframe,
	}

	err = db.InsertTrainData(newTraining)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding training & workshop data"})
		return
	}

	c.Redirect(http.StatusFound, "/training")
}
