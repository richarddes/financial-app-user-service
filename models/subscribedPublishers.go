package models

import (
	"context"
)

// SubscribedPublishers return the ids of the publishers the user with the id uid is subscribed to.
func (db *DB) SubscribedPublishers(ctx context.Context, uid uint64) ([]string, error) {
	query := `SELECT DISTINCT (unnest) FROM (
		SELECT unnest(subscribedPublisherIDs) FROM users WHERE id = $1
	) x;`

	rows, err := db.QueryContext(ctx, query, uid)
	if err != nil {
		return []string{}, err
	}

	defer rows.Close()

	var publisherIDs []string

	for rows.Next() {
		var publisherID string

		if err = rows.Scan(&publisherID); err != nil {
			return []string{}, err
		}

		publisherIDs = append(publisherIDs, publisherID)
	}

	return publisherIDs, nil
}
