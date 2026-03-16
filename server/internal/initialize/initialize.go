package initialize

import (
	"time"

	"github.com/InkSha/Resummer/internal/config"
	"github.com/InkSha/Resummer/internal/globals"
	"github.com/InkSha/Resummer/internal/service"
	"github.com/InkSha/Resummer/pkg/logger"
)

func setupLocation() {
	if loc, err := time.LoadLocation(config.Get("TIMEZONE")); err == nil {
		time.Local = loc
	} else {
		panic(err)
	}
}

func init() {
	config.LoadConfig()
}

func Initialize() {
	setupLocation()

	logger.SetupLogger(
		config.Get("LOG_LEVEL"),
		config.Get("LOG_OUTPUT"),
	)

	if dsn, err := config.GetDatabaseDSN(); err == nil {
		if db, err := initDB(dsn); err == nil {
			globals.SetDB(db)
			initModels(db, service.ListModels()...)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
