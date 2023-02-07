package main

import (
	"image/color"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/carepollo/alphacloud/getters"
	"github.com/carepollo/alphacloud/models"
	"github.com/fogleman/gg"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

func getDirname() string {
	_, filename, _, _ := runtime.Caller(0)
	currentFilePath, _ := filepath.Abs(filename)
	separated := strings.Split(currentFilePath, "/")
	separated = separated[:len(separated)-1]
	result := strings.Join(separated, "/")

	return result
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		var font, background string
		var color color.Color

		root := getDirname()
		filename := uuid.NewV4().String() + ".png"

		text := ctx.Query("text")
		if len(text) == 0 {
			return ctx.SendFile(root + "/public/error.jpg")
		}

		background = root + "/assets" + getters.GetBackground(ctx.Query("background"))
		color = getters.GetColor(ctx.Query("text_color"))
		font = root + "/assets" + getters.GetFont(ctx.Query("text_font"))
		requested := models.Request{
			Body:       text,
			Background: background,
			TextColor:  color,
			TextFont:   font,
		}

		picture, err := getters.Text2Image(requested)
		if err != nil {
			return ctx.SendFile(root + "/public/error.jpg")
		}

		err = gg.SavePNG("./public/"+filename, picture)
		if err != nil {
			return ctx.SendFile(root + "/public/error.jpg")
		}

		return ctx.SendFile(root + "/public/" + filename)
	})

	port := ":" + os.Getenv("PORT")
	app.Listen(port)
}
