package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



type apiConfig struct{
	database *database.Queries
}

func main(){
	godotenv.Load("../.env")
	urlDB := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", urlDB)
	if err != nil{
		log.Printf("Error when trying to open database: %w", err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{database: dbQueries}

	multiplexer := http.ServeMux{}
	multiplexer.HandleFunc("GET /v1/health", GetServerHealth)

	// Coupons api
	multiplexer.HandleFunc("POST /v1/coupons", apiCfg.HandleNewCoupon)
	multiplexer.HandleFunc("GET /v1/coupons", apiCfg.HandleGetCoupons)

	serverPort := os.Getenv("SERVER_PORT")
	server:= http.Server{Addr: fmt.Sprintf(":%v",serverPort), Handler: &multiplexer}
	server.ListenAndServe()
}


func GetServerHealth(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

