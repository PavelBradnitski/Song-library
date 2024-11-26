package main

import (
	"Songs/Song-library/internal/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
