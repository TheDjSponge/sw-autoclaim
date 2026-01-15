package scheduler

import (
	"context"
	"fmt"
	"time"
)


type Scheduler struct{
	claimInterval time.Duration
	cleanInterval time.Duration
	onClaim func()
	onClean func()
}

func NewScheduler(claimCouponInterval time.Duration, cleanInterval time.Duration, claimFunc, cleanFunc func()) *Scheduler{
	return &Scheduler{
		claimInterval: claimCouponInterval, 
		cleanInterval: cleanInterval,
		onClaim: claimFunc,
		onClean: cleanFunc,
	}
}


func (s *Scheduler) ScheduledTasksHandler(ctx context.Context){
	claimCouponsTicker := time.NewTicker(s.claimInterval)
	defer claimCouponsTicker.Stop()
	cleanCouponsTicker := time.NewTicker(s.cleanInterval)
	defer cleanCouponsTicker.Stop()
	

	fmt.Println("Starting initial coupon claim run")
	s.onClaim()
	for {
		select{
		case <- claimCouponsTicker.C:
			fmt.Println("Trying to fetch entries for redemption")
			s.onClaim()
		case <- cleanCouponsTicker.C:
			fmt.Println("Running coupon cleaning job")
			s.onClean()
		case <- ctx.Done():
			return
		}
	}
}
