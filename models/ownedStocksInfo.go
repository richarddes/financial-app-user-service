package models

import (
	"context"
	"user-service/config"
)

// OwnedStocksInfo returns the owned stocks a user with the specified uid has.
func (db *DB) OwnedStocksInfo(ctx context.Context, uid uint64) ([]config.StockInfo, error) {
	stmt := `
		SELECT (unnest).symbol, (unnest).amount, (unnest).boughtFor
		FROM (
			SELECT unnest(ownedStocks) FROM users WHERE id=$1
		) x;
	`

	rows, err := db.QueryContext(ctx, stmt, uid)
	if err != nil {
		return []config.StockInfo{}, err
	}

	shares := make([]config.StockInfo, 0)

	defer rows.Close()

	for rows.Next() {
		var (
			symbol    string
			amount    uint
			boughtFor float32
		)

		if err = rows.Scan(&symbol, &amount, &boughtFor); err != nil {
			return []config.StockInfo{}, err
		}

		shares = append(shares, config.StockInfo{Symbol: symbol, Amount: amount, BoughtFor: boughtFor})
	}

	return shares, nil
}
