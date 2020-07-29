// Package models implements functions manipulating the database.
package models

import (
	"context"
	"database/sql"
	"time"

	"user-service/config"
)

// DB represents a database connection.
type DB struct {
	*sql.DB
}

var retryCount = 1

// New returns a new instance of the DB struct. It tries to connect to
// the DB 5 times by waiting 10s after each failed attempt.
// If it fails 5 times to connect it returns an error.
func New(dbPath string) (*DB, error) {
	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second * 10)
	if err = db.Ping(); err != nil {
		if retryCount <= 5 {
			retryCount++
			New(dbPath)
		}

		return nil, err
	}

	return &DB{db}, nil
}

func (db *DB) stockInfo(ctx context.Context, uid uint64, body config.StockBody) (uint, float32, error) {
	var (
		ownedStocks uint
		boughtFor   float32
	)

	stmt := `
		SELECT (unnest).amount, (unnest).boughtFor
		FROM (
			SELECT unnest(ownedStocks) FROM users WHERE id=$1
		) x
		WHERE (unnest).symbol=$2;
	`

	err := db.QueryRowContext(ctx, stmt, uid, body.Symbol).Scan(&ownedStocks, &boughtFor)
	if err != nil {
		return 0, 0, err
	}

	return ownedStocks, boughtFor, nil
}
