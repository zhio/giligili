package server

import (
	"giligili/api"
	"giligili/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed := r.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.GET("user/me", api.UserMe)
			authed.DELETE("user/logout", api.UserLogout)
		}
		v1.POST("videos",api.CreateVideo)//restful api
		v1.GET("video/:id",api.ShowVideo)
		v1.GET("videos",api.ListVideo)
		v1.PUT("video/:id",api.UpdateVideo)
		v1.DELETE("videos/:id",api.DeleteVideo)

		v1.POST("upload/token",api.UploadToken)

		v1.GET("movies",api.ListMovies)
		v1.GET("movie/:id",api.ShowMovie)

	}
	return r
}
