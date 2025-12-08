package main

import (
	"encoding/json"
	"net/http"
	"time"

	"log"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)

const (
	CouponStatusNew = "new"
	CouponStatusActive = "active"
	CouponStatusExpired = "expired"
)

type DisplayableCouponInfo struct{
	Code        string
	Status      string
	FirstSeenAt time.Time
	UpdatedAt   time.Time
}

func ConvertToDisplayableCoupon(dbCoupon database.Coupon)(DisplayableCouponInfo){
	return DisplayableCouponInfo{
		Code: dbCoupon.Code,
		Status: dbCoupon.Status,
		FirstSeenAt: dbCoupon.FirstSeenAt,
		UpdatedAt: dbCoupon.UpdatedAt,
	}
}

func ConvertAllCouponsToDisplayable(dbCoupons []database.Coupon)([]DisplayableCouponInfo){
	displayableCoupons := make([]DisplayableCouponInfo, len(dbCoupons))
	for idx, coupon := range dbCoupons{
		displayableCoupons[idx] = ConvertToDisplayableCoupon(coupon)
	}
	return displayableCoupons
}


func (cfg *apiConfig) HandleNewCoupon(w http.ResponseWriter, r *http.Request){
	type bodyParams struct{
		Codes []string `json:"coupon_codes"`
	}

	var params bodyParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil{
		log.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "couldn't decode request parameters, check format.")
		return
	}

	if len(params.Codes) == 0{
		log.Printf("Received empty code slice.")
		respondWithError(w, http.StatusBadRequest, "no codes provided in request, check documentation for proper request format.")
		return
	}

	log.Printf("Trying to add the following codes to database: %v", params.Codes)
	for _, code := range params.Codes{
		err := cfg.database.AddCoupon(r.Context(),code)
		if err != nil{
			log.Println(err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	
	respondWithJSON(w, http.StatusCreated, JsonResponse{Status: http.StatusCreated, Message: "Inserted successfully"})
}

func (cfg *apiConfig) HandleGetCoupons(w http.ResponseWriter, r *http.Request){
	couponCode := r.URL.Query().Get("code")
	if couponCode == ""{
		coupons, err := cfg.database.GetAllCoupons(r.Context())
		if err != nil{
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve coupons")
			return
		}
		respondWithJSON(w, http.StatusOK, ConvertAllCouponsToDisplayable(coupons))
	}else{
		coupon, err := cfg.database.GetCouponByCode(r.Context(),couponCode)
		if err != nil{
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve coupon")
			return
		}
		respondWithJSON(w, http.StatusOK, ConvertToDisplayableCoupon(coupon))
	}
}
