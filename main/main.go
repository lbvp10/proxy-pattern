package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"proxy-pattern/libreria"
)

func main() {

	InitCache()

	app := fiber.New()

	app.Get("/:id<int>", getByIdHandler)
	app.Get("/metrics", getMetrics)

	log.Fatalln(app.Listen(":8089"))

}
func getByIdHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	var youtubeClient libreria.YoutubeClient = newVideoProxy()

	video := youtubeClient.GetVideo(id)

	return ctx.JSON(video)
}

func getMetrics(ctx *fiber.Ctx) error {
	return ctx.JSON(GetMetricGet())
}
