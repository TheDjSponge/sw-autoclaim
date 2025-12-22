package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	database          *database.Queries
	checkUserAPIURL   string
	claimCouponAPIURL string
}

func main() {
	godotenv.Load("../.env")
	urlDB := os.Getenv("DB_URL")
	checkUserAPIURL := os.Getenv("CHECK_USER_URL")
	claimCouponAPIURL := os.Getenv("CLAIM_COUPON_URL")

	db, err := pgx.Connect(context.Background(),urlDB)
	if err != nil {
		log.Printf("Error when trying to open database: %w", err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{database: dbQueries, checkUserAPIURL: checkUserAPIURL, claimCouponAPIURL: claimCouponAPIURL}

	stopHandler := make(chan struct{})

	go apiCfg.claimCouponRoutine(stopHandler)

	multiplexer := http.ServeMux{}
	multiplexer.HandleFunc("GET /v1/health", GetServerHealth)

	// Coupons api
	multiplexer.HandleFunc("POST /v1/coupons", apiCfg.HandleNewCoupon)
	multiplexer.HandleFunc("GET /v1/coupons", apiCfg.HandleGetCoupons)

	multiplexer.HandleFunc("POST /v1/users", apiCfg.HandleNewUser)
	multiplexer.HandleFunc("DELETE /v1/users", apiCfg.HandleDeleteUser)
	multiplexer.HandleFunc("GET /v1/users", apiCfg.HandleGetAllUsers)

	serverPort := os.Getenv("SERVER_PORT")
	server := http.Server{Addr: fmt.Sprintf(":%v", serverPort), Handler: &multiplexer}
	log.Println("Starting HTTP server")
	server.ListenAndServe()

	close(stopHandler)
}

func GetServerHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
