package routers

import "github.com/gin-gonic/gin"

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api/v1")
	CreateCoinRouter(apiV1)
	AuthenticationRouter(apiV1)
	return router
}