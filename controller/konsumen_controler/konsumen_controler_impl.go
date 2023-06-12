package konsumen_controller

import (
	"context"
	"test_kredit_plus/dto"
	"test_kredit_plus/helper"
	"test_kredit_plus/service/konsumen_service"
	"time"

	"github.com/gin-gonic/gin"
)

type KonsumenControllerImple struct {
	KonsumenService konsumen_service.KonsumenService
}

func NewKonsumenController(konsumenService konsumen_service.KonsumenService) KonsumenControler {
	return &KonsumenControllerImple{
		KonsumenService: konsumenService,
	}
}

func (controller *KonsumenControllerImple) Register(c *gin.Context) {
	var payloadKonsumen dto.RegisterKonsumen
	err := c.ShouldBindJSON(&payloadKonsumen)
	if err != nil {
		c.JSON(400, gin.H{
			"error_validasi": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	registerRespon, err := controller.KonsumenService.Register(ctx, payloadKonsumen)

	if err != nil {
		helper.SendInternalServerError(c, err)
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    registerRespon,
	})
}

func (conroller *KonsumenControllerImple) GetUser(c *gin.Context) {
	id := c.Value("user").(map[string]interface{})["id"].(string)

	detailKonsumen, err := conroller.KonsumenService.DetailKonsumen(c, id)
	if err != nil {
		helper.SendInternalServerError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"data":    detailKonsumen,
	})
}
