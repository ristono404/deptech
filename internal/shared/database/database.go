package database

import (
	pkgMySQL "github.com/ristono404/deptech/internal/pkg/database/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New(database pkgMySQL.Database) (*Database, error) {
	db, err := gorm.Open(mysql.Open(database.Source + "?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if database.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(database.MaxIdleConns)
	}
	if database.MaxOpenConns > 0 {
		sqlDB.SetMaxIdleConns(database.MaxOpenConns)
	}
	if database.ConnMaxLifetime.Nanoseconds() > 0 {
		sqlDB.SetConnMaxLifetime(database.ConnMaxLifetime)
	}

	return &Database{db}, nil
}
