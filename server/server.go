package server

import (
	"github.com/gofiber/fiber/v2"
	"proxy-pattern/libreria"
	"proxy-pattern/metrica"

	"proxy-pattern/proxy"
)

func AddRutas(app *fiber.App) {
	app.Get("/:id<int>", getByIdHandler)
	app.Get("/metrics", getMetricsHandler)
}

func getByIdHandler(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	var youtubeClient libreria.YoutubeClient = proxy.NewVideoProxy()

	video := youtubeClient.GetVideo(id)

	return ctx.JSON(video)
}

func getMetricsHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(metrica.GetMetricGet())
}
