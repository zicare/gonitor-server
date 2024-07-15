package controller

import (
	"github.com/gin-gonic/gin"
)

// StoryController exported
type StoryController struct {
}

// Fetch exported
func (sc StoryController) Fetch(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Story Fetch from user",
	})
}
