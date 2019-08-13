package api

import (
	"giligili/serializer"
	"github.com/gin-gonic/gin"
)

func CreateVideo(c *gin.Context){
	c.JSON(200,serializer.Response{
		Status: 0,
		Msg:    "成功",
	})
}