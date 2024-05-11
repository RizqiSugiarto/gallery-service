package main

import (
	"github.com/RizqiSugiarto/GaleryService/config"
	"github.com/RizqiSugiarto/GaleryService/internal/app"
)

func main() {
	cfg := config.Load()
	app.Run(cfg)
}