package models

import (
	"context"
	"user-service/config"
)

// SellStock sells body.Amount stocks of body.Symbol. If the user tries to sell more stocks
// than they own, the function will return a config.ErrBadRequest.
func (db *DB) SellStock(ctx context.Context, uid uint64, body config.StockBody) error {
	ownedStocks, boughtFor, err := db.stockInfo(ctx, uid, body)
	if err != nil {
		return err
	}

	var (
		stmt          string
		sharesLeft    = int(ownedStocks) - int(body.Amount)
		sharePrice    = float32(body.Amount) * body.Price
		totalShareVal = boughtFor - sharePrice
	)

	if sharesLeft > 0 {
		stmt = `
				UPDATE users 
				SET	cash = cash + $1, 
						ownedStocks[array_position(ownedStocks, (unnest)::stock)] = ($2, $3, $4)::stock
				FROM (
					SELECT (unnest) from (
						SELECT unnest(ownedStocks) FROM users WHERE id=$5
					) x WHERE (unnest).symbol = $2
				) y;
			`
	} else if sharesLeft == 0 {
		stmt = `
				UPDATE users 
				SET cash = cash + $1, 
						ownedStocks = array_remove(ownedStocks, ($2, $3, $4)::stock)
				WHERE id=$5;
			`

		// set sharesLeft & totalShareVal to the amount in the db to remove the entry
		sharesLeft = int(ownedStocks)
		totalShareVal = boughtFor
	} else {
		return config.ErrBadRequest
	}

	_, err = db.ExecContext(ctx, stmt, sharePrice, body.Symbol, sharesLeft, totalShareVal, uid)
	if err != nil {
		return err
	}

	return nil
}
