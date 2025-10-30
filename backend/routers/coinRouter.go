package routers

import (
	"github.com/VarisNithiparkorn/cryptoGraph/backend/controllers"
	"github.com/gin-gonic/gin"
)

func CreateCoinRouter(rg *gin.RouterGroup){
	coinRouter := rg.Group("/coins")

	coinRouter.GET("/",controllers.HandleGetAllCoins)
	coinRouter.GET("/ws", controllers.HandleWebSocketPrices)
}