package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"unico/internal/core/services"
	"unico/internal/handlers"
	"unico/internal/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()
	marketsRepository := repository.NewMarketRepositoryDb(dbClient)
	marketsService := services.New(marketsRepository)
	marketsHandler := handlers.NewHTTPHandler(marketsService)

	// define routes
	router.
		HandleFunc("/market/{searchType}/{id}", marketsHandler.Get).
		Methods(http.MethodGet).
		Name("GetMarket")
	router.
		HandleFunc("/market", marketsHandler.Post).
		Methods(http.MethodPost).
		Name("PostMarket")
	router.
		HandleFunc("/market/{id}", marketsHandler.Delete).
		Methods(http.MethodDelete).
		Name("DeleteMarket")
	router.
		HandleFunc("/market/{id}", marketsHandler.Put).
		Methods(http.MethodPut).
		Name("PutMarket")

		// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Print(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
