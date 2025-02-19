package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/vietddude/tss-impl/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once   sync.Once
)

// InitDB initializes the database connection pool
func InitDB(pgConfig *config.DB) (*pgxpool.Pool, error) {
	var err error
	var connStr string
	once.Do(func() {
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)
		dbPool, err = pgxpool.New(context.Background(), connStr)
		if err != nil {
			log.Printf("Unable to create connection pool: %v\n", err)
		}
	})

	return dbPool, err
}

// GetDB returns the database connection pool
func GetDB() *pgxpool.Pool {
	return dbPool
}

// CloseDB closes the database connection pool
func CloseDB() {
	if dbPool != nil {
		dbPool.Close()
	}
}
