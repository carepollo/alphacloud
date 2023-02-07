# alphacloud

this project is aiming to create a CDN to deploy on cloud. It is for academic purposes only.

## How to use

### Requirements
For better experience you need to have installed
- Git
- Go 1.20

### deployment

clone repo
`git clone https://github.com/carepollo/alphacloud.git`

create `.env` file with the following data
- `PORT` = [desired port]
- `DB_STRING_CONNECTION` = [your expected mongodb string connection]

then execute program
`go run main.go`

then go to your browser and visit the http://localhost:yourport

### API

You will request to the root path with the following query params
- `text`: mandatory param, is the text to be printed
- `background`: the background image to use, accepts the values `code`, `instagram`, `quote`, `white` default is `white`
- `text_color`: the color for the text `red`, `white`, `green`, `black` default is `black`
- `text_font`: the font for the text displayed, accepts `inconsolata`, `roboto`, `sassyfrass`, default is `roboto`
