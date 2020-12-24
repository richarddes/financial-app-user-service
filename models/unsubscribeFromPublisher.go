package models

import (
	"context"
	"errors"
)

// UnsubscribeFromPublisher lets a user with the id uid unsubscribe from a publisher with the id publisherID.
func (db *DB) UnsubscribeFromPublisher(ctx context.Context, uid uint64, publisherID string) error {
	if publisherID == "" {
		return errors.New("No publisherID has been specified")
	}

	stmt := `UPDATE users SET subscribedPublisherIDs = array_remove(subscribedPublisherIDs, $1) WHERE id = $2;`

	_, err := db.ExecContext(ctx, stmt, publisherID, uid)
	if err != nil {
		return err
	}

	return nil
}
