package service

import(
	"context"
	"giligili/serializer"
	"github.com/google/uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"time"
)

type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

func (service *UploadTokenService) Post() serializer.Response{
	u,_ := url.Parse(os.Getenv("COS_AD"))
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: os.Getenv("COS_SECRETID"),
			SecretKey: os.Getenv("COS_SECRETKEY"),
		},
	})


	key := "upload/avatar/" + uuid.Must(uuid.NewRandom()).String() + ".png"
	ctx := context.Background()
	presignedURL,err := c.Object.GetPresignedURL(ctx,http.MethodPut,key,os.Getenv("COS_SECRETID"),os.Getenv("COS_SECRETKEY"),time.Hour,nil)
	if err != nil{
		return serializer.Response{
			Status: 50002,
			Data:   nil,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	getpresignedURL, err := c.Object.GetPresignedURL(ctx, http.MethodGet, key,os.Getenv("COS_SECRETID"),os.Getenv("COS_SECRETKEY"), time.Hour, nil)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data:   map[string]string{
			"key": key,
			"put": presignedURL.String(),
			"get": getpresignedURL.String(),
		},

	}
}