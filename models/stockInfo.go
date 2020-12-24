package models

import (
	"context"
	"user-service/config"
)

// StockInfo returns the amount and total price for which a user with the id uid has bought a shares with the symbol body.Symbol.
func (db *DB) StockInfo(ctx context.Context, uid uint64, body config.StockBody) (uint, float32, error) {
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
