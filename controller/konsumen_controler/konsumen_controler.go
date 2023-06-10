package konsumen_controller

import "github.com/gin-gonic/gin"

type KonsumenControler interface {
	Register(ctx *gin.Context)
}
