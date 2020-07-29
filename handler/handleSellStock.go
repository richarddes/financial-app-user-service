package handler

import (
	"log"
	"net/http"
	"strconv"
	"user-service/config"
	"user-service/internal"
)

// HandleSellStock handles a user selling a stock.
// If the request's missing a UID header it will return a http.StatusBadRequest (http 400).
func HandleSellStock(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uidStr := r.Header.Get("UID")
		if uidStr == "" {
			http.Error(w, "A UID header's missing", http.StatusBadRequest)
			return
		}

		uid, err := strconv.ParseUint(uidStr, 10, 64)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var body config.StockBody

		err = internal.ParseJSONBody(r.Body, &body)
		if err != nil {
			http.Error(w, "Invalid request syntax", http.StatusBadRequest)
			log.Println(err)
			return
		}

		err = env.DB.SellStock(r.Context(), uid, body)
		if err != nil {
			if err == config.ErrBadRequest {
				http.Error(w, "Invalid request syntax", http.StatusBadRequest)
			} else {
				http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
				log.Println(err)
			}

			return
		}
	}
}
