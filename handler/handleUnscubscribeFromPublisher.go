package handler

import (
	"log"
	"net/http"
	"strconv"
	"user-service/config"
	"user-service/internal"
)

// HandleUnsubscribeFromPublisher lets a user with the id uid unsubscribe from a publisher with an id specified in the request body.
func HandleUnsubscribeFromPublisher(env *config.Env) http.HandlerFunc {
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

		var body config.SubscribeBody

		err = internal.ParseJSONBody(r.Body, &body)
		if err != nil {
			http.Error(w, "Invalid request syntax", http.StatusBadRequest)
			return
		}

		err = env.DB.UnsubscribeFromPublisher(r.Context(), uid, body.PublisherID)
		if err != nil {
			http.Error(w, "An unexpected error has occured while unsubscribing from the publisher", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}
