package api

import (
	"giligili/service/movie"
	"github.com/gin-gonic/gin"
)

func ShowMovie(c *gin.Context) {
	service := movie.ShowMovieService{}
	var res= service.Show(c.Param("id"))
	c.JSON(200, res)
}
//视频列表接口
func ListMovies(c *gin.Context){
	service := movie.ListMovieService{}
	var res = service.List()
	c.JSON(200,res)
}
