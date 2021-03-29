package banking

import (
	"github.com/alekssro/banking/banking/shared/logger"
	"github.com/alekssro/banking/cmd/banking/app"
)

func main() {
	logger.Info("Banking server listening...")
	app.Start()
}
