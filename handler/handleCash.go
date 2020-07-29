package handler

import (
	"log"
	"net/http"
	"strconv"
	"user-service/config"
)

// HandleCash returns the amount of cash a users has.
// If the request's missing a UID header it will return a http.StatusBadRequest (http 400).
func HandleCash(env *config.Env) http.HandlerFunc {
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

		cash, err := env.DB.Cash(r.Context(), uid)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Write([]byte(strconv.FormatFloat(float64(cash), 'f', 2, 32)))
	}
}
