package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	libreria "proxy-pattern/libreria"
	"sync"
	"time"
)

var cache_ *cache.Cache
var once sync.Once

const metric_get = "metric_get"

type VideoProxy struct {
	Video     *libreria.Video `json:"video"`
	idsBanned []int
}
type VideoMetric struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}

func newVideoProxy() *VideoProxy {
	return &VideoProxy{
		Video:     &libreria.Video{},
		idsBanned: []int{3, 4, 7, 8, 9, 10},
	}
}

func (videoProxy VideoProxy) GetVideo(id int) *libreria.Video {
	fmt.Printf("Buscando video don id:%d\n", id)
	idString := string(id)
	if v, found := cache_.Get(idString); found {
		videoProxy.Video = v.(*libreria.Video)
		fmt.Printf("Video encontrado en chache id:%d\n", videoProxy.Video.Id)
	} else {
		videoProxy.Video = videoProxy.Video.GetVideo(id)
		fmt.Printf("Video No encontrado en chache id:%d\n", videoProxy.Video.Id)
		cache_.Add(idString, videoProxy.Video, cache.DefaultExpiration)
	}
	sendMetricGet()
	return videoProxy.Video
}

func (videoProxy VideoProxy) PostVideo(video *libreria.Video) *libreria.Video {
	return video.PostVideo(video)
}

func InitCache() *cache.Cache {
	once.Do(func() { //Forzar a que este bloque de codigo se ejecute una unica vez en la app
		fmt.Println("Inicializando cache.......")
		cache_ = cache.New(5*time.Minute, 10*time.Minute)
		fmt.Println("Inicializando cache OK.......")
	})
	cache_.Add(metric_get, 0, cache.NoExpiration)
	return cache_
}
func sendMetricGet() {
	cache_.Increment(metric_get, 1)
}

func GetMetricGet() *VideoMetric {
	count, _ := cache_.Get(metric_get)
	return &VideoMetric{
		Id:    metric_get,
		Count: count.(int),
	}

}
