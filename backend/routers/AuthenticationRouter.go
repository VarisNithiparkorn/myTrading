package routers

import (
	"github.com/VarisNithiparkorn/cryptoGraph/controllers"
	"github.com/gin-gonic/gin"
)

func AuthenticationRouter(rg *gin.RouterGroup) {
	rg.POST("/login",controllers.HandleLogin)
	authRouter := rg.Group("/register")
	authRouter.POST("",controllers.HandleRegister)
	authRouter.GET("",controllers.HandleVerifyEmail)
}
