package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "alive"})
}
