package driver

import (
	"errors"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/purwalenta/purwalenta/pkg/config"
)

var (
	once              sync.Once
	postgreDBInstance *sqlx.DB
	postgreDBError    error

	ErrEmptyDSNString = errors.New("dsn string could not be empty")
)

const (
	postgreDriverName = "postgres"
)

func GetPostgreDriver() (*sqlx.DB, error) {
	once.Do(func() {
		cfg := config.GetConfig().Database.Postgre

		if cfg.DSN == "" {
			postgreDBInstance, postgreDBError = nil, ErrEmptyDSNString
			return
		}

		postgreDBInstance, postgreDBError = sqlx.Open(postgreDriverName, cfg.DSN)
		if nil != postgreDBError {
			return
		}

		postgreDBInstance.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)
		postgreDBInstance.SetMaxIdleConns(cfg.MaxIdle)
		postgreDBInstance.SetMaxOpenConns(cfg.MaxOpen)
	})

	return postgreDBInstance, postgreDBError
}
