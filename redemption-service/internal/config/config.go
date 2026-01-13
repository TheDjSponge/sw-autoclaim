package config

import (
	"os"

	"github.com/joho/godotenv"
)

type APIConfig struct {
	DBConnURL             string
	CheckUserAPIURL   string
	ClaimCouponAPIURL string
}

func LoadConfig() APIConfig{
	godotenv.Load("../.env")
	urlDB := os.Getenv("DB_URL")
	checkUserAPIURL := os.Getenv("CHECK_USER_URL")
	claimCouponAPIURL := os.Getenv("CLAIM_COUPON_URL")
	return APIConfig{DBConnURL: urlDB, CheckUserAPIURL: checkUserAPIURL, ClaimCouponAPIURL: claimCouponAPIURL}
}