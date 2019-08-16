package model

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Video struct {
	gorm.Model
	Title 	string
	Info string
	URL string
	Avatar string
}
//封面地址
func (video *Video) AvatarURL() string{
	u,_ := url.Parse(os.Getenv("COS_AD"))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
		},
	})
	ctx := context.Background()
	presignedURL,_ := c.Object.GetPresignedURL(ctx,http.MethodGet,video.Avatar,os.Getenv("COS_SECRETID"),os.Getenv("COS_SECRETKEY"),time.Hour,nil)
	return presignedURL.String()
}