package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/api"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/config"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/coupons"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/redemption"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/scheduler"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/users"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const CouponClaimInterval = time.Hour * 24
const CouponCleanInterval = time.Hour * 24 * 7

func main() {
	cfg := config.LoadConfig()

	dbConfig, _ := pgxpool.ParseConfig(cfg.DBConnURL)
	dbConfig.MaxConns = 10
	dbConfig.MinConns = 1
	dbConfig.MaxConnLifetime = 1 * time.Hour
	dbConfig.MaxConnIdleTime = 30 * time.Minute
	dbConfig.HealthCheckPeriod = 5 * time.Minute
	
	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Printf("Error when trying to open database: %v", err)
	}
	defer pool.Close()
	dbQueries := database.New(pool)

	userValidator := users.HiveValidator{CheckUserURL: cfg.CheckUserAPIURL}
	userService := users.NewService(dbQueries, userValidator)
	couponService := coupons.NewService(dbQueries)
	redemptionService := redemption.NewService(dbQueries, cfg.ClaimCouponAPIURL)
	scheduler := scheduler.NewScheduler(
		CouponClaimInterval, 
		CouponCleanInterval,
		func() {redemptionService.ClaimNewRedemptions()},
		func() {couponService.CleanExpiredCoupons(context.Background())},
	)
	handler := api.NewHandler(userService, couponService, redemptionService)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go scheduler.ScheduledTasksHandler(ctx)
	
	multiplexer := http.ServeMux{}
	handler.InitRoutes(&multiplexer)
	serverPort := os.Getenv("SERVER_PORT")
	server := http.Server{Addr: fmt.Sprintf(":%v", serverPort), Handler: &multiplexer}
	log.Println("Starting HTTP server")
	server.ListenAndServe()
}

