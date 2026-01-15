package redemption

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)

type Service struct{
	db *database.Queries
	claimCouponAPIURL string
}

func NewService(db *database.Queries, claimCouponAPIURL string) *Service{
	return &Service{db:db, claimCouponAPIURL: claimCouponAPIURL}
}

type APICode int
const (
	APICodeSuccess APICode = 100
	APICodeCouponInvalid APICode = 302
	APICodeCouponAlreadyClaimed APICode = 304
	APICodeCouponExpired APICode = 306
)

const (
	RedemptionSuccess = "success"
	RedemptionFailed = "failed"
	RedemptionAlreadyClaimed = "already_claimed"
)

type RetCode int;

func (r *RetCode) UnmarshalJSON(b []byte) error {
	// Custom unmarshalling method to handle retCodes returned as integers (e.g 100 for successful) or strings (e.g "(H304)" for error codes)
	var s string
	err := json.Unmarshal(b, &s)
	if err == nil{
		code, err := ExtractRetCodeFromString(s)
		if err != nil{
			return err
		}
		*r = RetCode(code)
		return nil
	}

	var i int
	err = json.Unmarshal(b, &i)
	if err == nil{
		*r = RetCode(i)
		return nil
	}

	return errors.New("unexpected retcode format received")
}

func ExtractRetCodeFromString(s string) (int, error){
	// Extracts an n digit code formatted as "H(xyz)" where xyz is an arbitrary code defined by the api endpoint.
	if len(s) < 2 {
		return -1, errors.New("got unexpected RetCode string")
	}
	code, err := strconv.Atoi(s[2:len(s)-1])
	if err != nil{
		return -1, err
	}
	return code, nil
}



func (s *Service) ClaimCoupon(hive_id string, server string, code string) (APICode, string, error) {
	formData := url.Values{
		"country": {"CH"},
		"lang":    {"fr"},
		"server":  {server},
		"hiveid":  {hive_id},
		"coupon":  {code},
	}
	encodedData := formData.Encode()

	req, err := http.NewRequest("POST", s.claimCouponAPIURL, strings.NewReader(encodedData))
	if err != nil {
		fmt.Println("throwing at request creation")
		return -1, "", err
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
		RetCode  RetCode         `json:"retCode"`
		RetMsg   string          `json:"retMsg"`
		UserData userDataPayload `json:"userData"`
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("throwing at request")
		return -1, "", err
	}
	defer resp.Body.Close()
	var respPayload responsePayload
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&respPayload)
	if err != nil {
		fmt.Println("throwing at decoding")
		return -1, "", err
	}
	
	return APICode(respPayload.RetCode), respPayload.RetMsg, nil
}

func (s *Service) ClaimCouponsForUser(ctx context.Context, hive_id string) error {
	coupons, err := s.db.GetRedemptionsForUser(ctx, hive_id)
	redemptionEntries := make([]database.AddRedemptionsParams, 0,len(coupons))
	if err != nil{
		return fmt.Errorf("failed to get redemptions entries for user: %w", err)
	}
	for _, redemption := range coupons{
		entry, err := s.processRedemption(context.Background(), database.GetUnclaimedRedemptionsRow{
			ID: redemption.UserID,
			HiveID: redemption.HiveID,
			Server: redemption.Server,
			ID_2: redemption.CouponID,
			Code: redemption.Code,
		})
		if err != nil{
			fmt.Printf("Error when trying to claim coupon %v for user %v, and err : %v", redemption.Code, redemption.HiveID, err)
		}
		redemptionEntries = append(redemptionEntries, entry)
		time.Sleep(2 * time.Second)
	}

	_, err = s.db.AddRedemptions(context.Background(), redemptionEntries)
	if err != nil{
		return fmt.Errorf("failed to insert new redemptions : %w", err)
	}
	return nil
}

func (s *Service) ClaimNewRedemptions() error {
	unclaimed, err := s.db.GetUnclaimedRedemptions(context.Background())
	if err != nil{
		return fmt.Errorf("error while fetching unclaimed redemptions %w", err)
	}
	redemptionEntries := make([]database.AddRedemptionsParams, 0,len(unclaimed))
	
	for _, redemption := range unclaimed{
		entry, err := s.processRedemption(context.Background(),redemption)
		if err != nil{
			fmt.Printf("Error when trying to claim coupon %v for user %v, and err : %v", redemption.Code, redemption.HiveID, err)
		}
		redemptionEntries = append(redemptionEntries, entry)
		time.Sleep(2 * time.Second)
	}
	
	_, err = s.db.AddRedemptions(context.Background(), redemptionEntries)
	if err != nil{
		return fmt.Errorf("failed to insert new redemptions : %w", err)
	}
	return nil
}


func (s *Service) processRedemption(ctx context.Context, redemption database.GetUnclaimedRedemptionsRow) (database.AddRedemptionsParams, error){
	// Processes one redemption job. A redemption job consists in a coupon-user pair involving any user and a new coupon. 
	code, msg, err := s.ClaimCoupon(redemption.HiveID, redemption.Server, redemption.Code)
	if err != nil{
		fmt.Printf("Error when trying to claim coupon %v for user %v, got message : %v , and err : %v", redemption.Code, redemption.HiveID, msg, err)
	}

	switch code{
	case APICodeSuccess:
		_ = s.db.SetCouponActive(ctx, redemption.ID_2)
		log.Printf("Successfully claimed coupon: [%s] for profile: [%s / %s]",redemption.Code, redemption.HiveID, redemption.Server)
		return successRedemption(redemption, msg), nil
	
	case APICodeCouponInvalid:
		_ = s.db.SetCouponExpired(ctx, redemption.ID_2)
		return failedRedemption(redemption, msg), nil

	case APICodeCouponExpired:
		_ = s.db.SetCouponExpired(ctx, redemption.ID_2)
		return failedRedemption(redemption, msg), nil
	
	case APICodeCouponAlreadyClaimed:
		_ = s.db.SetCouponActive(ctx, redemption.ID_2)
		return alreadyClaimedRedemption(redemption, msg), nil
	default:
		return failedRedemption(redemption, msg), fmt.Errorf("couldn't handle retCode: [%v] properly", code)
	}

}

func successRedemption(r database.GetUnclaimedRedemptionsRow, msg string) database.AddRedemptionsParams {
	return database.AddRedemptionsParams{
		UserID: r.ID,
		CouponID: r.ID_2,
		Status: RedemptionSuccess, 
		ResponseMessage: msg,
	}
}

func failedRedemption(r database.GetUnclaimedRedemptionsRow, msg string) database.AddRedemptionsParams {
	return database.AddRedemptionsParams{
		UserID: r.ID,
		CouponID: r.ID_2,
		Status: RedemptionFailed, 
		ResponseMessage: msg,
	}
}

func alreadyClaimedRedemption(r database.GetUnclaimedRedemptionsRow, msg string) database.AddRedemptionsParams {
	return database.AddRedemptionsParams{
		UserID: r.ID,
		CouponID: r.ID_2,
		Status: RedemptionAlreadyClaimed, 
		ResponseMessage: msg,
	}
}
