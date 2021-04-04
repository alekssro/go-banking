package main

import (
	"github.com/alekssro/banking/banking/api"
	"github.com/alekssro/banking/banking/shared/logger"
)

func main() {
	logger.Info("Banking server listening...")
	api.Start()
}
