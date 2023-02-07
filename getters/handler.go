package getters

import (
	"image"

	"github.com/carepollo/alphacloud/models"
	"github.com/fogleman/gg"
)

func Text2Image(config models.Request) (image.Image, error) {
	bgImage, err := gg.LoadImage(config.Background)
	if err != nil {
		return nil, err
	}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	result := gg.NewContext(imgWidth, imgHeight)
	result.DrawImage(bgImage, 0, 0)

	if err := result.LoadFontFace(config.TextFont, 80.0); err != nil {
		return nil, err
	}

	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - 80)
	maxWidth := float64(imgWidth) - 60.0

	result.SetColor(config.TextColor)
	result.DrawStringWrapped(config.Body, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return result.Image(), nil
}
