package handler_test

import (
	"net/http"
	"testing"
	"user-service/config"
	"user-service/handler"
)

func TestHandleBuyStock(t *testing.T) {
	mockEnv := config.NewMockEnv()

	cases := []struct {
		body         config.StockBody
		expectedCode int
	}{
		{config.StockBody{Symbol: "AAPL", Amount: 5, Price: 250.3}, http.StatusBadRequest},
		{config.StockBody{Symbol: "FB", Amount: 10, Price: 143.74}, http.StatusBadRequest},
		{config.StockBody{Symbol: "AAPL", Amount: 2, Price: 250.3}, http.StatusOK},
		{config.StockBody{Symbol: "FB", Amount: 5, Price: 143.74}, http.StatusOK},
	}

	for _, c := range cases {
		rr := handler.NewRecorder(t, "POST", "/api/users/buy-stock", c.body, handler.HandleBuyStock(mockEnv))

		if rr.Code != c.expectedCode {
			t.Errorf("Expected status code %d but got %d instead when body=%v", c.expectedCode, rr.Code, c.body)
		}
	}
}
