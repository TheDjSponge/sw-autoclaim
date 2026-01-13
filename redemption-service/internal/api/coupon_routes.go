package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)

type DisplayableCouponInfo struct {
	Code        string
	Status      string
	FirstSeenAt time.Time
	UpdatedAt   time.Time
}

func ConvertToDisplayableCoupon(dbCoupon database.Coupon) DisplayableCouponInfo {
	return DisplayableCouponInfo{
		Code:        dbCoupon.Code,
		Status:      dbCoupon.Status,
		FirstSeenAt: dbCoupon.FirstSeenAt.Time,
		UpdatedAt:   dbCoupon.UpdatedAt.Time,
	}
}

func ConvertAllCouponsToDisplayable(dbCoupons []database.Coupon) []DisplayableCouponInfo {
	displayableCoupons := make([]DisplayableCouponInfo, len(dbCoupons))
	for idx, coupon := range dbCoupons {
		displayableCoupons[idx] = ConvertToDisplayableCoupon(coupon)
	}
	return displayableCoupons
}



func (h *Handler) HandleNewCoupon(w http.ResponseWriter, r *http.Request) {
	type bodyParams struct {
		Codes []string `json:"coupon_codes"`
	}

	var params bodyParams
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Println(err.Error())
		RespondWithMessage(w, http.StatusBadRequest, "couldn't decode request parameters, check format.")
		return
	}

	log.Printf("Trying to add the following codes to database: %v", params.Codes)
	err = h.couponService.AddCouponBatch(r.Context(), params.Codes)
	if err != nil{
		RespondWithMessage(w, http.StatusInternalServerError, err.Error())
	}

	RespondWithJSON(w, http.StatusCreated, JsonResponse{Status: http.StatusCreated, Message: "Inserted successfully"})
}

func (h *Handler) HandleGetCoupons(w http.ResponseWriter, r *http.Request) {
	couponCode := r.URL.Query().Get("code")
	if couponCode == "" {
		coupons, err := h.couponService.GetAllCoupons(r.Context())
		if err != nil {
			RespondWithMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
		RespondWithJSON(w, http.StatusOK, ConvertAllCouponsToDisplayable(coupons))
	} else {
		coupon, err := h.couponService.GetCouponByCode(r.Context(), couponCode)
		if err != nil {
			RespondWithMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
		RespondWithJSON(w, http.StatusOK, ConvertToDisplayableCoupon(coupon))
	}
}
