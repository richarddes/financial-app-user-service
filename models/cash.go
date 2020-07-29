package models

import "context"

// Cash returns the amount of cash a user with the specified uid has.
func (db *DB) Cash(ctx context.Context, uid uint64) (float32, error) {
	var cash float32

	qry := "SELECT cash from users WHERE id=$1;"

	err := db.QueryRowContext(ctx, qry, uid).Scan(&cash)
	if err != nil {
		return 0, err
	}

	return cash, nil
}
