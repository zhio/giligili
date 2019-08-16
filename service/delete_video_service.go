package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情服务
type DeleteVideoService struct {
}

// Show 展示视频
func (service *DeleteVideoService) Delete(id string) serializer.Response{
	var video model.Video
	err := model.DB.First(&video,id).Error
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error: err.Error(),
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil{
		return serializer.Response{
			Status: 50000,
			Data:   nil,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
	}
}

