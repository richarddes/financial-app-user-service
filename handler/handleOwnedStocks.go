package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user-service/config"
)

// HandleOwnedStocks returns the stocks owned by a user.
// If the request's missing a UID header it will return a http.StatusBadRequest (http 400).
func HandleOwnedStocks(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uidStr := r.Header.Get("UID")
		if uidStr == "" {
			http.Error(w, "A UID header's missing", http.StatusBadRequest)
			return
		}

		uid, err := strconv.ParseUint(uidStr, 10, 64)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		ownedShares, err := env.DB.OwnedStocksInfo(r.Context(), uid)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		// The output json, must always have a data field when stock data is being send.
		outFormat := make(map[string][]config.StockInfo, 1)

		outFormat["data"] = ownedShares

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(outFormat)
	}
}
