package handler

import (
	"log"
	"net/http"
	"strconv"
	"user-service/config"
	"user-service/internal"
)

// HandleChangeLang changes the language of a user.
// If the request's missing a UID or Lang header it will return a http.StatusBadRequest (http 400).
func HandleChangeLang(env *config.Env) http.HandlerFunc {
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

		lang := r.Header.Get("Lang")
		if !internal.IsSupportedLang(lang) {
			http.Error(w, "The specified language isn't a supported language", http.StatusBadRequest)
			return
		}

		err = env.DB.SetLang(r.Context(), uid, lang)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
