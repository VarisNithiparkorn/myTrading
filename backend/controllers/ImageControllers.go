package controllers

import "github.com/gin-gonic/gin"

func HandleGetImageByName(c *gin.Context) {
	name := c.Query("name")
	c.File("template/images/"+name )
}