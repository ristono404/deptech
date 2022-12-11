package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Source          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func New(database Database) (*sql.DB, error) {
	db, err := sql.Open("mysql", database.Source+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	if database.MaxIdleConns > 0 {
		db.SetMaxIdleConns(database.MaxIdleConns)
	}
	if database.MaxOpenConns > 0 {
		db.SetMaxOpenConns(database.MaxOpenConns)
	}
	if database.ConnMaxLifetime.Nanoseconds() > 0 {
		db.SetConnMaxLifetime(database.ConnMaxLifetime)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
