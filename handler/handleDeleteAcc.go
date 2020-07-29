package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"
	"user-service/config"
)

// HandleDeleteAcc deletes a user account.
// If the request's missing a UID header it will return a http.StatusBadRequest (http 400).
func HandleDeleteAcc(env *config.Env) http.HandlerFunc {
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

		err = env.DB.DeleteAcc(r.Context(), uid)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "auth_token",
			Value:    "",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
			Secure:   true,
		})
	}
}
