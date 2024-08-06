package employee

import (
	"log"
	"net/http"

	"job/db"

	"github.com/gin-gonic/gin"
)

func LanguageHandler(c *gin.Context) {

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

	language := c.PostForm("language")
	speak := c.PostForm("speak")
	read := c.PostForm("read")
	write := c.PostForm("write")

	newLang := db.Language{
		Lang:  language,
		Speak: speak,
		Read:  read,
		Write: write,
	}

	err = db.InsertLanguage(newLang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error adding language data"})
		return
	}

	c.Redirect(http.StatusFound, "/language")
}
