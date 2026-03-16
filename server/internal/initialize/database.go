package initialize

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, &InitDBError{msg: "Init DB Failed: dsn is empty"}
	}

	return gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{},
	)
}

func initModels(db *gorm.DB, models ...any) {
	db.AutoMigrate(models...)
}
