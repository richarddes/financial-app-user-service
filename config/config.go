// Package config defines globally used interfaces and structs.
package config

import (
	"context"
	"errors"
)

var (
	// ErrBadRequest defines an error which triggers a StatusBadRequest (http 400) to be sent
	ErrBadRequest error = errors.New("The user's syntax is invalid")

	// SupportedLangs defines the languages supported by the proxied services.
	// It should be set once the program starts.
	SupportedLangs []string
)

type (
	// Env represents a collection of interfaces required for the handlers.
	Env struct {
		DB Datastore
	}

	// SubscribeBody represents the body needed to (un)subscribe to a new publisher
	SubscribeBody struct {
		PublisherID string `json:"publisherID"`
	}

	// StockBody represents the expeceted body when a stock bought or sold
	StockBody struct {
		Symbol string  `json:"symbol"`
		Amount uint    `json:"amount"`
		Price  float32 `json:"price"`
	}

	// StockInfo represents the returned data from a bought stocks
	StockInfo struct {
		Symbol    string  `json:"symbol"`
		Amount    uint    `json:"amount"`
		BoughtFor float32 `json:"boughtFor"`
	}
)

type (
	// Datastore defines functions a datastore has to implement.
	Datastore interface {
		DeleteAcc(ctx context.Context, uid uint64) error
		Cash(ctx context.Context, uid uint64) (float32, error)
		BuyStock(ctx context.Context, uid uint64, body StockBody) error
		SellStock(ctx context.Context, uid uint64, body StockBody) error
		OwnedStocksInfo(ctx context.Context, uid uint64) ([]StockInfo, error)
		SetLang(ctx context.Context, uid uint64, lang string) error
		SubscribeToPublisher(ctx context.Context, uid uint64, publisherID string) error
		UnsubscribeFromPublisher(ctx context.Context, uid uint64, publisherID string) error
		SubscribedPublishers(ctx context.Context, uid uint64) ([]string, error)
	}
)
