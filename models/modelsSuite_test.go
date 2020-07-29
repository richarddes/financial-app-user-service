// +build integration

package models_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"user-service/config"
	"user-service/models"

	_ "github.com/lib/pq"
)

const (
	cash float32 = 1000
)

var (
	uid uint64
	db  *sql.DB

	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")

	stocks = []struct {
		body  config.StockBody
		valid bool
	}{
		{config.StockBody{Symbol: "AAPL", Amount: 500, Price: 200.00}, false},
		{config.StockBody{Symbol: "AAPL", Amount: 5, Price: 200.00}, true},
	}
	validStocks = 1
)

func init() {
	if dbUser == "" {
		log.Fatal("No environment variable named DB_USER present")
	}

	if dbPass == "" {
		log.Fatal("No environment variable named DB_PASSWORD present")
	}

	if dbPort == "" {
		log.Fatal("No environment variable named DB_PORT present")
	}

	if dbName == "" {
		log.Fatal("No environment variable named DB_NAME present")
	}

	if dbHost == "" {
		dbHost = "localhost"
	}

	// test values
	config.SupportedLangs = []string{"en", "de"}
}

func TestDefaultImpl(t *testing.T) {
	connStr := fmt.Sprintf("port=%s user=%s password=%s dbname=%s host=%s sslmode=disable", dbPort, dbUser, dbPass, dbName, dbHost)
	database, err := models.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	db = database.DB

	setupDB(t, database.DB)

	ModelsSuite(t, database)
}

func ModelsSuite(t *testing.T, impl config.Datastore) {
	ctx := context.Background()

	t.Run("test language change", func(t *testing.T) {
		langs := []struct {
			lang  string
			valid bool
		}{
			{"fr", false},
			{"us", false},
			{"en", true},
			{"de", true},
		}

		for _, l := range langs {
			err := impl.SetLang(ctx, uid, l.lang)
			if err != nil && l.valid {
				t.Errorf("Unexpected error: %v when language=%s", err, l.lang)
			} else if err == nil && !l.valid {
				t.Errorf("Expected an error but got none when language=%s", l.lang)
			}
		}
	})

	t.Run("test that cash() returns the correct amount", func(t *testing.T) {
		c, err := impl.Cash(ctx, uid)
		if err != nil {
			t.Fatal(err)
		}

		if c != cash {
			t.Fatal("The cash amount in the database doesn't equal the specified amount")
		}
	})

	t.Run("test stock buying", func(t *testing.T) {

		for _, s := range stocks {
			err := impl.BuyStock(ctx, uid, s.body)
			if err != nil && s.valid {
				t.Errorf("Unexpected error: %v when body=%v", err, s.body)
			} else if err == nil && !s.valid {
				t.Errorf("Expected an err but got none when body=%v", s.body)
			}

			if s.valid {
				c, err := impl.Cash(ctx, uid)
				if err != nil {
					t.Fatal(err)
				}

				price := float32(s.body.Amount) * s.body.Price
				expectedCashLeft := cash - price

				if c != expectedCashLeft {
					t.Errorf("Expected %f to be left when buying shares for a total of %f", expectedCashLeft, price)
				}
			}
		}
	})

	t.Run("test stock info retrieval", func(t *testing.T) {
		shs, err := impl.OwnedStocksInfo(ctx, uid)
		if err != nil {
			t.Fatal(err)
		}

		if len(shs) != validStocks {
			t.Errorf("Expected %d stocks to be bought but got %d instead", validStocks, len(shs))
		}
	})

	t.Run("test stock selling", func(t *testing.T) {
		for _, s := range stocks {
			err := impl.SellStock(ctx, uid, s.body)
			if err != nil && s.valid {
				if err != nil && s.valid {
					t.Errorf("Unexpected error: %v when body=%v", err, s.body)
				} else if err == nil && !s.valid {
					t.Errorf("Expected an err but got none when body=%v", s.body)
				}
			}

			shs, err := impl.OwnedStocksInfo(ctx, uid)
			if err != nil {
				t.Fatal(err)
			}

			if containsStock(t, shs, s.body) {
				t.Errorf("The stock hasn't been sold when the stock=%v", s)
			}
		}
	})

	t.Run("test account deletion", func(t *testing.T) {
		err := impl.DeleteAcc(ctx, uid)
		if err != nil {
			t.Fatal(err)
		}

		if models.UsersExists(db) {
			t.Errorf("The user entry hasn't been deleted if it should've been")
		}
	})
}

func setupDB(t *testing.T, db *sql.DB) {
	t.Helper()

	insertUser := `INSERT INTO users VALUES (DEFAULT,$1,$2,$3,$4,$5,$6,$7::stock[]) RETURNING id;`

	err := db.QueryRow(insertUser, "john.doe@email.com", "password", "doe", "john", "en", cash, "{}").Scan(&uid)
	if err != nil {
		t.Fatal(err)
	}
}

func containsStock(t *testing.T, stocks []config.StockInfo, stock config.StockBody) bool {
	t.Helper()

	for _, s := range stocks {
		if s.Symbol == stock.Symbol && s.Amount == stock.Amount {
			return true
		}
	}

	return false
}
