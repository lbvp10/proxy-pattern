package metrica

import (
	"github.com/patrickmn/go-cache"
	server "proxy-pattern/cache"
)

const Metric_get = "Metric_get"

type VideoMetric struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}

func InitMetrics() {
	server.DoCache().Add(Metric_get, 0, cache.NoExpiration)
}

func SendMetricGet() {
	server.DoCache().Increment(Metric_get, 1)
}

func GetMetricGet() *VideoMetric {
	count, _ := server.DoCache().Get(Metric_get)
	return &VideoMetric{
		Id:    Metric_get,
		Count: count.(int),
	}

}
