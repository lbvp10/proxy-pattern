package libreria

import (
	"crypto/rand"
	"math/big"
	"time"
)

type Video struct {
	Id     int
	Url    string
	Nombre string
	Data   string
}

type YoutubeClient interface {
	GetVideo(id int) *Video
	PostVideo(video *Video) *Video
}

func (v Video) GetVideo(id int) *Video {
	time.Sleep(5 * time.Second) // Simulando un delay por consumir otra API o una Query en db pesada 10seg
	var video Video
	if id%2 == 0 { // Construyendo la respuesta quemada
		video = Video{Id: id, Url: "https://www.youtube.com/watch?v=p7zevPed3Ss", Nombre: "¡Tengo miedo!", Data: ""}
	} else {
		video = Video{Id: id, Url: "https://www.youtube.com/watch?v=Y0az9jHNwRU", Nombre: "Ciclista sin casco", Data: ""}
	}
	return &video
}

func (v Video) PostVideo(video *Video) *Video {
	video.Id = generateId()
	time.Sleep(1 * time.Second) // Simulando un delay por consumir otra API o una Query en db pesada 10seg
	return video
}

func generateId() int {
	numero, _ := rand.Int(rand.Reader, big.NewInt(1000))
	return int(numero.Int64())
}
