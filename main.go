package main

import (
	"image/color"
	"log"
	"os"

	"github.com/carepollo/alphacloud/getters"
	"github.com/carepollo/alphacloud/models"
	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		var font, background string
		var color color.Color

		text := ctx.Query("text")
		backgroundString := ctx.Query("background")
		colorString := ctx.Query("text_color")
		fontString := ctx.Query("text_cont")

		background = getters.GetBackground(backgroundString)
		color = getters.GetColor(colorString)
		font = getters.GetFont(fontString)
		requested := models.Request{
			Body:       text,
			Background: background,
			TextColor:  color,
			TextFont:   font,
		}

		picture, err := Text2Image(requested)
		if err != nil {
			return ctx.SendString("internal server error: failed to generate image")
		}

		err = gg.SavePNG("./public/", picture)
		if err != nil {
			return ctx.SendString("internal server error: failed to save file")
		}

		return ctx.SendString("this is a test")
	})

	port := ":" + os.Getenv("PORT")
	app.Listen(port)
}
