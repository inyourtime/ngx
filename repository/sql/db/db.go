package db

import (
	"ngx/domain"
	"ngx/port"
	"ngx/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	config util.Config
	logger port.Logger
	gormDB *gorm.DB
}

func New(config util.Config, logger port.Logger) (*Database, error) {
	var err error
	database := &Database{
		config: config,
		logger: logger,
	}

	database.gormDB, err = database.connect()
	if err != nil {
		return nil, err
	}

	err = database.autoMigrate()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (db *Database) DB() *gorm.DB {
	return db.gormDB
}

func (db *Database) connect() (*gorm.DB, error) {
	gdb, err := gorm.Open(postgres.Open(db.config.PostgresSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		DryRun: false,
		// TranslateError: true,
	})
	if err != nil {
		return nil, err
	}

	return gdb, nil
}

func (db *Database) autoMigrate() error {
	return db.gormDB.AutoMigrate(&domain.User{})
}
