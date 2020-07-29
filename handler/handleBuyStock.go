// Package handler implements all http handlers.
package handler

import (
	"log"
	"net/http"
	"strconv"
	"user-service/config"
	"user-service/internal"
)

// HandleBuyStock a user buying a stock.
// If the request's missing a UID header it will return a http.StatusBadRequest (http 400).
func HandleBuyStock(env *config.Env) http.HandlerFunc {
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

		var body config.StockBody

		err = internal.ParseJSONBody(r.Body, &body)
		if err != nil {
			http.Error(w, "Invalid request syntax", http.StatusBadRequest)
			return
		}

		err = env.DB.BuyStock(r.Context(), uid, body)
		if err != nil {
			if err == config.ErrBadRequest {
				http.Error(w, "You don't have enough money to buy that many shares", http.StatusBadRequest)
			} else {
				http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
				log.Println(err)
			}

			return
		}
	}
}
