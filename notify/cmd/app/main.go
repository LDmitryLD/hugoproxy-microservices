package main

import (
	"os"
	"projects/LDmitryLD/hugoproxy-microservices/notify/config"
	"projects/LDmitryLD/hugoproxy-microservices/notify/internal/infrastructure/logs"
	"projects/LDmitryLD/hugoproxy-microservices/notify/run"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	conf := config.NewAppConf()

	logger := logs.NewLogger(conf, os.Stdout)

	conf.Init(logger)

	app := run.NewApp(conf, logger)

	if code := app.Run(); code != 0 {
		logger.Info("app run error")
		os.Exit(code)
	}
}
