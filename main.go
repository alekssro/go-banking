package main

import (
	"github.com/alekssro/banking/app"
	"github.com/alekssro/banking/logger"
)

func main() {
	logger.Info("Banking server listening...")
	app.Start()
}
