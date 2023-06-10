package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code":   500,
		"status": "INTERNAL SERVER ERROR",
		"error":  err.Error(),
	})
	c.Abort()
}
