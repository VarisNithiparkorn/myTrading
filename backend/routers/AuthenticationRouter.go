package routers

import (
	"github.com/VarisNithiparkorn/cryptoGraph/controllers"
	"github.com/gin-gonic/gin"
)

func AuthenticationRouter(rg *gin.RouterGroup) {
	authRouter := rg.Group("/register")
	authRouter.POST("",controllers.HandleRegister)
	authRouter.GET("",controllers.HandleVerifyEmail)
}
