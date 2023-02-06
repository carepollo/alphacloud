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

You will have the following endpoints

