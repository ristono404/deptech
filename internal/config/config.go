package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	pkgMySQL "github.com/ristono404/deptech/internal/pkg/database/mysql"
	"github.com/spf13/cast"
)

type Config struct {
	Database pkgMySQL.Database
	PerPage  uint64
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	config := &Config{
		Database: pkgMySQL.Database{
			Source:          os.Getenv("DATABASE"),
			MaxIdleConns:    cast.ToInt(os.Getenv("DATABASE_MAX_IDLE_CONNS")),
			MaxOpenConns:    cast.ToInt(os.Getenv("DATABASE_MAX_OPEN_CONNS")),
			ConnMaxLifetime: time.Duration(cast.ToInt(os.Getenv("DATABASE_CONN_MAX_LIFETIME_MINUTES"))) * time.Minute,
		},
		PerPage: cast.ToUint64(os.Getenv("PER_PAGE")),
	}

	return config
}
