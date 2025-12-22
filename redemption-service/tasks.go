package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)

const CouponClaimInterval = 24 * time.Hour // 3 days

func (cfg *apiConfig) claimCouponRoutine(stopHandler chan struct{}){
	ticker := time.NewTicker(CouponClaimInterval)
	defer ticker.Stop()

	fmt.Println("Starting initial coupon claim run")
	ClaimNewRedemptions(cfg)
	for {
		select{
		case <- ticker.C:
			fmt.Println("Trying to fetch entries for redemption")
			ClaimNewRedemptions(cfg)
		case <- stopHandler:
			return
		}
	}
}


func ClaimNewRedemptions(cfg *apiConfig) error {
	unclaimed, err := cfg.database.GetUnclaimedRedemptions(context.Background())
	redemptionEntries := make([]database.AddRedemptionsParams, 0,len(unclaimed))
	if err != nil{
		log.Printf("Error while trying to fetch unclaimed redemptions (coupon-user pairs) : %v", err)
		return err
	}
	for _, redemption := range unclaimed{
		fmt.Printf("Attempting to claim coupon : %v for hive_id : %v",redemption.Code, redemption.HiveID)
		isRedeemed, msg, err := cfg.ClaimCoupon(redemption.HiveID, redemption.Server, redemption.Code)
		if err != nil{
			fmt.Printf("Error when trying to claim coupon %v for user %v, got message %v and err %v", redemption.Code, redemption.HiveID, msg, err)
		}
		if isRedeemed{
			fmt.Println("Coupon claimed successfully!")
			redemptionEntries = append(redemptionEntries, database.AddRedemptionsParams{UserID: redemption.ID, CouponID: redemption.ID_2,Status: "success",ResponseMessage: msg})
		} else {
			fmt.Println("Coupon claiming failed!")
			redemptionEntries = append(redemptionEntries, database.AddRedemptionsParams{UserID: redemption.ID, CouponID: redemption.ID_2,Status: "failed", ResponseMessage: msg})
		}
		time.Sleep(2*time.Second)
	}
	
	fmt.Printf("Got redemption entries : %v", redemptionEntries)
	bla, err := cfg.database.AddRedemptions(context.Background(), redemptionEntries)
	fmt.Println(bla)
	if err != nil{
		fmt.Printf("Failed to add redemption entries to DB, got error : %v", err)
		return fmt.Errorf("Failed to add redemption entries to DB, got error : %v", err)
	}
	return nil
}



func (cfg *apiConfig) ClaimCoupon(hive_id string, server string, code string) (bool, string, error) {
	formData := url.Values{
		"country": {"CH"},
		"lang":    {"fr"},
		"server":  {server},
		"hiveid":  {hive_id},
		"coupon":  {code},
	}
	encodedData := formData.Encode()

	req, err := http.NewRequest("POST", cfg.claimCouponAPIURL, strings.NewReader(encodedData))
	if err != nil {
		fmt.Println("throwing at request creation")
		return false, "", err
	}

	req.Header.Set("Referer", "https://event.withhive.com/ci/smon/evt_coupon")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	type userDataPayload struct {
		UID        int    `json:"uid"`
		ServerName string `json:"server_name"`
		WizardName string `json:"wizard_name"`
	}

	type responsePayload struct {
		RetCode  any             `json:"retCode"`
		RetMsg   string          `json:"retMsg"`
		UserData userDataPayload `json:"userData"`
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("throwing at request")
		return false, "", err
	}
	defer resp.Body.Close()
	var respPayload responsePayload
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respPayload)
	if err != nil {
		fmt.Println("throwing at decoding")
		return false, "", err
	}
	// display response payload
	if respPayload.RetCode != "100" {
		return false, "", fmt.Errorf("check user endpoint hit failed with code: %v", respPayload)
	}
	return true, respPayload.RetMsg, nil
}