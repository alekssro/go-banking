package main

import (
	"github.com/alekssro/banking/auth/api"
	"github.com/alekssro/banking/lib/logger"
)

func main() {
	logger.Info("Starting Authentication server...")
	api.Start()
}
