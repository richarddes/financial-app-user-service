package models

import (
	"context"
	"errors"
	"user-service/internal"
)

// SetLang sets the user's language with the specified uid to the specified language.
func (db *DB) SetLang(ctx context.Context, uid uint64, lang string) error {
	if !internal.IsSupportedLang(lang) {
		return errors.New("The specified language is not supported")
	}

	stmt := "UPDATE users SET lang = $1 WHERE id = $2;"

	_, err := db.ExecContext(ctx, stmt, lang, uid)
	if err != nil {
		return err
	}

	return nil
}
