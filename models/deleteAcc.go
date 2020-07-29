package models

import "context"

// DeleteAcc deletes the user account with the specified uid,
func (db *DB) DeleteAcc(ctx context.Context, uid uint64) error {
	deleteAccStmt := "DELETE FROM users WHERE id=$1;"

	_, err := db.ExecContext(ctx, deleteAccStmt, uid)
	if err != nil {
		return err
	}

	return nil
}
