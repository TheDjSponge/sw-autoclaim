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
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

const CouponClaimInterval = time.Hour * 24
const CouponCleanInterval = time.Hour * 24 * 7

func main() {
	cfg := config.LoadConfig()

	db, err := pgx.Connect(context.Background(),cfg.DBConnURL)
	if err != nil {
		log.Printf("Error when trying to open database: %v", err)
	}
	dbQueries := database.New(db)

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

