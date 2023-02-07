package getters

// returns path of font expected, selects default if given data does not exist
func GetFont(param string) string {
	switch param {
	case "inconsolata":
		return "fonts/Inconsolata-Regular.ttf"
	case "roboto":
		return "/fonts/Roboto-Regular.ttf"
	case "sassyfrass":
		return "/fonts/SassyFrass-Regular.tts"
	default:
		return "/fonts/Roboto-Regular.ttf"
	}
}
