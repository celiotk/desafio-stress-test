package usecase

import (
	"context"
	"fmt"
	"stress_test/internal/entity"
	"sync/atomic"
	"time"
)

type StressTestInputDTO struct {
	URL         string
	Requests    int
	Concurrency int
}

type StressTestOutputDTO struct {
	TotalTime   Duration
	TotalReq    int
	Status200   int
	OtherStatus map[int]int
	Errors      int
}

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, time.Duration(d).String())), nil
}

type StressTestUsecase struct {
	requestService entity.ServiceRequestInterface
}

func NewStressTest(sr entity.ServiceRequestInterface) *StressTestUsecase {
	return &StressTestUsecase{sr}
}

func (st *StressTestUsecase) Run(ctx context.Context, input StressTestInputDTO) (report StressTestOutputDTO, err error) {
	chOut := make(chan workerResult)
	requests := int64(input.Requests)
	report.OtherStatus = make(map[int]int)

	inicio := time.Now()
	for i := 0; i < input.Concurrency; i++ {
		go st.worker(ctx, input.URL, &requests, chOut)
	}

	remaing := input.Requests
resLoop:
	for {
		select {
		case res := <-chOut:
			if res.err != nil {
				report.Errors++
			} else {
				report.TotalReq++
				if res.status == 200 {
					report.Status200++
				} else {
					report.OtherStatus[res.status]++
				}
			}
			remaing--
			if remaing == 0 {
				break resLoop
			}
		case <-ctx.Done():
			return report, ctx.Err()
		}
	}
	report.TotalTime = Duration(time.Since(inicio))
	return report, nil
}

type workerResult struct {
	status int
	err    error
}

func (st *StressTestUsecase) worker(ctx context.Context, url string, requests *int64, chOut chan workerResult) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		if atomic.AddInt64(requests, -1) < 0 {
			return
		}
		status, err := st.requestService.SendRequest(ctx, url)
		chOut <- workerResult{status, err}
	}
}
