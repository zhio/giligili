package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
)
//视频投稿接口
func CreateVideo(c *gin.Context){
	service := service.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil{
		res := service.Create()
		c.JSON(200,res)
	}else {
		c.JSON(200,ErrorResponse(err))
	}
}
//视频详情接口
func ShowVideo(c *gin.Context) {
	service := service.ShowVideoService{}
	var res= service.Show(c.Param("id"))
	c.JSON(200, res)
}
//视频列表接口
func ListVideo(c *gin.Context){
	service := service.ListVideoService{}
	var res = service.List()
	c.JSON(200,res)
}
//UpdateVideo 更新视频的接口
func UpdateVideo(c *gin.Context) {
	service := service.UpdateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		var res= service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
//DeleteVideo 删除视频接口
func DeleteVideo(c *gin.Context) {
	service := service.DeleteVideoService{}
	var res= service.Delete(c.Param("id"))
	c.JSON(200, res)

}