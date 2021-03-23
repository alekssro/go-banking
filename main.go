package main

import (
	"github.com/alekssro/banking/app"
	"github.com/alekssro/banking/logger"
)

func main() {
	logger.Info("Starting banking server...")
	app.Start()
}
