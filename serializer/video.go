package serializer

import "giligili/model"

type Video struct {
	ID        uint   `json:"id"`
	Title  string `json:"title"`
	Info  string `json:"info"`
	URL 	string 	`json:"url"`
	Avatar	string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
//创建视频的序列化器
func BuildVideo(item model.Video) Video{
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:	   item.URL,
		Avatar:    item.AvatarURL(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}
//序列化视频列表
func BuildVideos(items []model.Video) (videos []Video){
	for _,item := range items{
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}