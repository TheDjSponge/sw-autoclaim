package coupons

import (
	"context"
	"fmt"

	"github.com/TheDjSponge/sw-autoclaim/redemption-service/internal/database"
)


const (
	CouponStatusNew     = "new"
	CouponStatusActive  = "active"
	CouponStatusExpired = "expired"
)

type Service struct{
	db		*database.Queries
}

func NewService(db *database.Queries) *Service{
	return &Service{
		db: db,
	}
}

func (s *Service) GetCouponByCode(ctx context.Context, code string) (database.Coupon, error) {
	coupon, err := s.db.GetCouponByCode(ctx, code)
	if err != nil{
		return database.Coupon{}, fmt.Errorf("failed to retrieve coupon by code: %w", err)
	}
	return coupon, nil
}

func (s *Service) GetAllCoupons(ctx context.Context) ([]database.Coupon, error){
	coupons, err := s.db.GetAllCoupons(ctx)
	if err != nil{
		return []database.Coupon{}, fmt.Errorf("failed getting all coupons: %w", err)
	}
	return coupons, nil
}

func (s *Service) AddCoupon(ctx context.Context, code string) error {
	err := s.db.AddCoupon(ctx, code)
	if err != nil {
		return fmt.Errorf("failed adding coupon: %w", err)
	}
	return err
}

func (s *Service) AddCouponBatch(ctx context.Context, codes []string) error {
	if len(codes) == 0 {
		return fmt.Errorf("failed when adding batch coupons: No codes provided")
	}

	for _, code := range codes {
		err := s.AddCoupon(ctx, code)
		if err != nil {
			return fmt.Errorf("failed when adding batch coupons: %w", err)
		}
	}
	return nil
}