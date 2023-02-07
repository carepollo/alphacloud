package getters

import "image/color"

// returns expected color by given data, selects default if given data is not acceptable
func GetColor(param string) color.Color {
	var result color.Color

	switch param {
	case "black":
		result = color.Black
	case "white":
		result = color.White
	case "red":
		result = color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		}
	case "green":
		result = color.RGBA{
			R: 0,
			G: 128,
			B: 0,
			A: 255,
		}
	default:
		result = color.Black
	}

	return result
}
