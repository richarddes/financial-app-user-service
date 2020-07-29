package handler_test

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"user-service/config"
	"user-service/handler"
)

func TestHandleOwnedShares(t *testing.T) {
	mockEnv := config.NewMockEnv()

	// same content as returned by the mockEnv's db's OwnedSharesInfo function
	body := []config.StockInfo{
		{Symbol: "AAPL", Amount: 100, BoughtFor: 25034.00},
		{Symbol: "FB", Amount: 30, BoughtFor: 7530.50},
		{Symbol: "SNAP", Amount: 5, BoughtFor: 1000.00},
	}

	rr := handler.NewRecorder(t, "GET", "/api/users/owned-shares", nil, handler.HandleOwnedStocks(mockEnv))

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d instead", http.StatusOK, rr.Code)
	}

	var resp map[string][]config.StockInfo

	err := json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(body, resp["data"]) {
		t.Errorf("The returned owned stocks(=%v) don't eqaul the expected owned shares(=%v)", resp, body)
	}
}
