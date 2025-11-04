package routers

import (
	"github.com/VarisNithiparkorn/cryptoGraph/controllers"
	"github.com/gin-gonic/gin"
)

func ImageRouter(rg *gin.RouterGroup) {
	imageRouter := rg.Group("/images")
	imageRouter.GET("",controllers.HandleGetImageByName)
}