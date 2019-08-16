package movie

import (
	"fmt"
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情服务
type ShowMovieService struct {
}

// Show 展示视频
func (service *ShowMovieService) Show(id string) serializer.Response{
	var movie model.Movie
	err := model.DB.First(&movie,id).Error
	if err != nil{
		return serializer.Response{
			Status: 404,
			Msg:    "电影不存在",
			Error: err.Error(),
		}
	}
	fmt.Println(serializer.BuildMovie(movie))
	return serializer.Response{
		Msg: "get",
		Data:   serializer.BuildMovie(movie),
	}
}

