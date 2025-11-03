package routers

import (
	"github.com/VarisNithiparkorn/cryptoGraph/backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthenticationRouter(rg *gin.RouterGroup) {
	authRouter := rg.Group("/register")
	authRouter.POST("",controllers.HandleRegister)
}