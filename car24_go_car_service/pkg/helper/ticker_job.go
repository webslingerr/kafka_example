package helper

import (
	"fmt"
	"time"

	"gitlab.udevs.io/car24/car24_go_car_service/config"
)

type JobTicker struct {
	T *time.Timer
}

func GetNextTickDuration() time.Duration {
	now := time.Now()
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), config.HOUR_TO_TICK, config.MINUTE_TO_TICK, config.SECOND_TO_TICK, 0, time.Local)
	if nextTick.Before(now) {
		nextTick = nextTick.Add(config.INTERVAL_PERIOD)
	}
	return nextTick.Sub(time.Now())
}

func NewJobTicker() JobTicker {
	fmt.Println("New Tick Here....")
	return JobTicker{time.NewTimer(GetNextTickDuration())}
}

func (jt JobTicker) UpdateJobTicker() {
	fmt.Println("NEXT Tick Here....")
	jt.T.Reset(GetNextTickDuration())
}
