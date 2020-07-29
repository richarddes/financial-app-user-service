package config

import (
	"context"
)

type mockDB struct {
	//ownedStocks represents a map of company tags and the amount of shares owned
	ownedStock map[string]uint
}

func (db *mockDB) DeleteAcc(ctx context.Context, uid uint64) error {
	return nil
}

func (db *mockDB) Cash(ctx context.Context, uid uint64) (float32, error) {
	return 1000.00, nil
}

func (db *mockDB) BuyStock(ctx context.Context, uid uint64, body StockBody) error {
	total := float32(body.Amount) * body.Price
	cash, _ := db.Cash(ctx, 1)

	if cash-total < 0 {
		return ErrBadRequest
	}

	return nil
}

func (db *mockDB) SellStock(ctx context.Context, uid uint64, body StockBody) error {
	if body.Amount > db.ownedStock[body.Symbol] {
		return ErrBadRequest
	}

	return nil
}

func (db *mockDB) OwnedStocksInfo(ctx context.Context, uid uint64) ([]StockInfo, error) {
	return []StockInfo{
		{Symbol: "AAPL", Amount: 100, BoughtFor: 25034.00},
		{Symbol: "FB", Amount: 30, BoughtFor: 7530.50},
		{Symbol: "SNAP", Amount: 5, BoughtFor: 1000.00},
	}, nil
}

func (db *mockDB) SetLang(ctx context.Context, uid uint64, lang string) error {
	return nil
}

// NewMockEnv returns a new Env with mock values instead of production values.
func NewMockEnv() *Env {
	env := new(Env)

	db := new(mockDB)
	db.ownedStock = make(map[string]uint)

	db.ownedStock["AAPL"] = 10
	db.ownedStock["FB"] = 30

	env.DB = db

	return env
}
