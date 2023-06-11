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

func SendStatusOk(c *gin.Context, message string, data map[string]interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    data,
		"status":  "status Ok",
		"message": message,
	})
	c.Abort()
}
