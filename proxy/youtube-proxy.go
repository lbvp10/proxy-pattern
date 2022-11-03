package proxy

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	cache2 "proxy-pattern/cache"
	libreria "proxy-pattern/libreria"
	"proxy-pattern/metrica"
)

type VideoProxy struct {
	Video     *libreria.Video `json:"video"`
	idsBanned []int
}

func NewVideoProxy() *VideoProxy {
	return &VideoProxy{
		Video:     &libreria.Video{},
		idsBanned: []int{3, 4, 7, 8, 9, 10},
	}
}

func (videoProxy VideoProxy) GetVideo(id int) *libreria.Video {
	fmt.Printf("Buscando video don id:%d\n", id)
	//Validar los banneds aca
	idString := string(id)
	if v, found := cache2.DoCache().Get(idString); found {
		videoProxy.Video = v.(*libreria.Video)
		fmt.Printf("Video encontrado en chache id:%d\n", videoProxy.Video.Id)
	} else {
		videoProxy.Video = videoProxy.Video.GetVideo(id)
		fmt.Printf("Video No encontrado en chache id:%d\n", videoProxy.Video.Id)
		cache2.DoCache().Add(idString, videoProxy.Video, cache.DefaultExpiration)
	}
	metrica.SendMetricGet()
	return videoProxy.Video
}

func (videoProxy VideoProxy) PostVideo(video *libreria.Video) *libreria.Video {
	return video.PostVideo(video)
}
