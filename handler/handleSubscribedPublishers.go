package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user-service/config"
)

// HandleSubscribedPublishers returns the publisher ids of the publishers a user with the id uid is subscribed to.
func HandleSubscribedPublishers(env *config.Env) http.HandlerFunc {
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

		publisherIds, err := env.DB.SubscribedPublishers(r.Context(), uid)
		if err != nil {
			http.Error(w, "An unexpected error has occured while retrieving the subscribed publishers", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(publisherIds)
	}
}
