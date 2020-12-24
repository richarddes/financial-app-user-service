package models

import (
	"context"
	"errors"
)

// SubscribeToPublisher lets the user with the id uid subscribe to the publisher with the id publisherID.
func (db *DB) SubscribeToPublisher(ctx context.Context, uid uint64, publisherID string) error {
	if publisherID == "" {
		return errors.New("No publisherID has been specified")
	}

	stmt := "UPDATE users SET subscribedPublisherIDs = array_append(subscribedPublisherIDs, $1) WHERE id = $2;"

	_, err := db.ExecContext(ctx, stmt, publisherID, uid)
	if err != nil {
		return err
	}

	return nil
}
