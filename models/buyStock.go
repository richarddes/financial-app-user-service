package models

import (
	"context"
	"database/sql"
	"user-service/config"
)

// BuyStock first checks if the user has enough money to buy body.Amount many shares. 
// If so the users owned stocks entry is updated, but if not the function returns config.ErrBadRequest.
func (db *DB) BuyStock(ctx context.Context, uid uint64, body config.StockBody) error {
	var (
		stmt         string
		insertAmount = body.Amount
		sharePrice   = float32(body.Amount) * body.Price
	)

	cash, err := db.Cash(ctx, uid)
	if err != nil {
		return err
	}

	if cash-sharePrice < 0 {
		return config.ErrBadRequest
	}

	ownedStocks, boughtFor, err := db.stockInfo(ctx, uid, body)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		stmt = `
			UPDATE users
			SET cash = cash - $1, ownedStocks = ownedStocks || ARRAY[($2, $3, $4)::stock]
			WHERE id=$5;
		`
	} else {
		stmt = `
			UPDATE users 
			SET	cash = cash - $1, ownedStocks[array_position(ownedStocks, (unnest)::stock)] = ($2, $3, $4)::stock
			FROM (
				SELECT (unnest) from (
					SELECT unnest(ownedStocks) FROM users WHERE id=$5
				) x WHERE (unnest).symbol = $2
			) y;
		`
		insertAmount += ownedStocks
	}

	totalShareVal := boughtFor + sharePrice

	_, err = db.ExecContext(ctx, stmt, sharePrice, body.Symbol, insertAmount, totalShareVal, uid)
	if err != nil {
		return err
	}

	return nil
}
