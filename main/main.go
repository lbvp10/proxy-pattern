package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"proxy-pattern/cache"
	"proxy-pattern/metrica"
	"proxy-pattern/server"
)

func main() {

	cache.DoCache()
	metrica.InitMetrics()

	app := fiber.New()

	server.AddRutas(app)

	log.Fatalln(app.Listen(":8089"))

}
