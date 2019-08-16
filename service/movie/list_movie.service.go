package movie

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情服务
type ListMovieService struct {
}


// Show 展示全部视频
func (service *ListMovieService) List() serializer.Response{
	var movies []model.Movie
	err := model.DB.Limit(24).Find(&movies).Error
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Msg:    "数据库查询错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data:   serializer.BuildMovies(movies),
	}
}