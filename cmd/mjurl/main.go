package main

import (
	"github.com/jcserv/mjurl/internal"
	"github.com/jcserv/mjurl/internal/utils/log"
	"go.uber.org/zap"
)

func main() {
	logger := log.Init()
	defer logger.Sync()

	service, err := internal.NewMJURLService()
	if err != nil {
		logger.Fatal("could not create service", zap.Error(err))
	}

	defer service.Shutdown()

	if err := service.Run(); err != nil {
		logger.Fatal("could not start service", zap.Error(err))
	}
}
