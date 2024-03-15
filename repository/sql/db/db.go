package db

import (
	"database/sql"
	"ngx/port"
	"ngx/util"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type Database struct {
	config util.Config
	logger port.Logger
	bunDB  *bun.DB
}

func New(config util.Config, logger port.Logger) (*Database, error) {
	var err error
	database := &Database{
		config: config,
		logger: logger,
	}

	database.bunDB, err = database.connect()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (db *Database) DB() *bun.DB {
	return db.bunDB
}

func (db *Database) connect() (*bun.DB, error) {
	dsn := db.config.PostgresSource
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	bunDB := bun.NewDB(pgdb, pgdialect.New())

	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	err := bunDB.Ping()
	if err != nil {
		return nil, err
	}

	return bunDB, nil
}
