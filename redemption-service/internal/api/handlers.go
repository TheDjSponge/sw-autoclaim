package api

import (
	"net/http"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/coupons"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/redemption"
	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/users"
)

type Handler struct{
	redemptionService *redemption.Service
	userService *users.Service
	couponService *coupons.Service
}

func NewHandler(userService *users.Service, couponService *coupons.Service, redemptionService *redemption.Service) *Handler{
	return &Handler{
		userService: userService, 
		couponService: couponService,
		redemptionService: redemptionService,
	}
}

func (h *Handler) InitRoutes(mux *http.ServeMux){
	mux.HandleFunc("GET /v1/health", GetServerHealth)

	// Coupons api
	mux.HandleFunc("POST /v1/coupons", h.HandleNewCoupon)
	mux.HandleFunc("GET /v1/coupons", h.HandleGetCoupons)

	// Users api
	mux.HandleFunc("POST /v1/users", h.HandleNewUser)
	mux.HandleFunc("DELETE /v1/users", h.HandleDeleteUser)
	mux.HandleFunc("GET /v1/users", h.HandleGetAllUsers)
}