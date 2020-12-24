package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"user-service/config"
	"user-service/models"

	"github.com/gorilla/mux"

	"user-service/handler"

	_ "github.com/lib/pq"
)

var (
	dbUser   = os.Getenv("DB_USER")
	dbPass   = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
	dbName   = os.Getenv("DB_NAME")
	dbHost   = os.Getenv("DB_HOST")
	devMode  = os.Getenv("DEV_MODE")
	sptLangs = os.Getenv("SUPPORTED_LANGUAGES")
)

func init() {
	if dbUser == "" {
		log.Fatal("No environment variable named DB_USER present")
	}

	if dbPass == "" {
		log.Fatal("No environment variable named DB_PASSWORD present")
	}

	if dbPort == "" {
		log.Fatal("No environment variable named DB_PORT present")
	}

	if dbName == "" {
		log.Fatal("No environment variable named DB_NAME present")
	}

	if dbHost == "" {
		dbHost = "localhost"
	}

	if devMode == "" {
		devMode = "true"
	}

	if sptLangs == "" {
		config.SupportedLangs = []string{"en"}
	}

	config.SupportedLangs = strings.Split(sptLangs, ",")
}

func main() {
	connStr := fmt.Sprintf("port=%s user=%s password=%s dbname=%s host=%s sslmode=disable", dbPort, dbUser, dbPass, dbName, dbHost)
	db, err := models.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	env := &config.Env{DB: db}

	r := mux.NewRouter()

	api := r.PathPrefix("/api/users").Subrouter()
	api.HandleFunc("/change-lang", handler.HandleChangeLang(env)).Methods("POST")
	api.HandleFunc("/delete-acc", handler.HandleDeleteAcc(env)).Methods("POST")
	api.HandleFunc("/sell-stock", handler.HandleSellStock(env)).Methods("POST")
	api.HandleFunc("/buy-stock", handler.HandleBuyStock(env)).Methods("POST")
	api.HandleFunc("/subscribe-to-publisher", handler.HandleSubscribeToPublisher(env)).Methods("POST")
	api.HandleFunc("/unsubscribe-from-publisher", handler.HandleUnsubscribeFromPublisher(env)).Methods("POST")
	api.HandleFunc("/subscribed-publishers", handler.HandleSubscribedPublishers(env)).Methods("GET")
	api.HandleFunc("/cash", handler.HandleCash(env)).Methods("GET")
	api.HandleFunc("/owned-stocks", handler.HandleOwnedStocks(env)).Methods("GET")

	fmt.Println("The user-service is ready")
	log.Fatal(http.ListenAndServe(":8081", r))
}
