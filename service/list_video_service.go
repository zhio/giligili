package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情服务
type ListVideoService struct {
}


// Show 展示全部视频
func (service *ListVideoService) List() serializer.Response{
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Msg:    "数据库查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:   serializer.BuildVideos(videos),
	}
}