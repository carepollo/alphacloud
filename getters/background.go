package getters

// returns path of background file selected, selects default given data does not exist
func GetBackground(param string) string {
	switch param {
	case "code":
		return "/backgrounds/bg_code.jpg"
	case "instagram":
		return "/backgrounds/bg_instagram.jpg"
	case "quote":
		return "/backgrounds/bg_quote.jpg"
	default:
		return "/backgrounds/bg_white.jpg"
	}
}
