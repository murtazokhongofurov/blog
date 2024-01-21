package main

import (
	"github.com/blog/config"
	"github.com/blog/internal/app"
)

func main() {
	cfg := config.Load()
	app.Run(cfg)
}
