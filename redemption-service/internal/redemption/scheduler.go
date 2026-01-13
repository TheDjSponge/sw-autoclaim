package redemption

import (
	"context"
	"fmt"
	"time"
)


type Scheduler struct{
	redemptionService *Service
	claimInterval time.Duration
}

func NewScheduler(redemptionService *Service, claimCouponInterval time.Duration) *Scheduler{
	return &Scheduler{redemptionService: redemptionService, claimInterval: claimCouponInterval}
}


func (s *Scheduler) ScheduledTasksHandler(ctx context.Context){
	claimCouponsTicker := time.NewTicker(s.claimInterval)
	defer claimCouponsTicker.Stop()

	fmt.Println("Starting initial coupon claim run")
	s.redemptionService.ClaimNewRedemptions()
	for {
		select{
		case <- claimCouponsTicker.C:
			fmt.Println("Trying to fetch entries for redemption")
			s.redemptionService.ClaimNewRedemptions()
		case <- ctx.Done():
			return
		}
	}
}
