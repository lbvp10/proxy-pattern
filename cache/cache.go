package cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var cache_ *cache.Cache
var once sync.Once

func DoCache() *cache.Cache {
	once.Do(func() { //Forzar a que este bloque de codigo se ejecute una unica vez en la app
		fmt.Println("Inicializando cache.......")
		cache_ = cache.New(5*time.Minute, 10*time.Minute)
		fmt.Println("Inicializando cache OK.......")
	})
	return cache_
}
