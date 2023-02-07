package models

import "image/color"

type Request struct {
	/*
		the background image, it can be 'code', 'instagram', 'quote', 'white', by default must be 'white'
	*/
	Background string

	/*
		color of text, it can be 'red', 'white', 'green', 'black' default is 'black'
	*/
	TextColor color.Color

	/*
		font of text, can be 'inconsolata', 'roboto', 'sassyfrass', default is 'roboto'
	*/
	TextFont string

	/*
		actual message to display
	*/
	Body string
}
