package lib

import (
	"fmt"
	"math"
	"sync"
)

type Result struct {
	maxLatency float64
	minLatency float64
	avgLatency float64
	noOfRequestCompleted int
	noOfSuccessfulRequest int
	mutex *sync.Mutex
}

func NewResult() *Result {
	return &Result{
		mutex: &sync.Mutex{},
		minLatency: math.MaxFloat64,
	}
}

func (res *Result) AddResult(latency float64, isSuccess bool) {
	res.mutex.Lock()
	defer res.mutex.Unlock()
	res.maxLatency = math.Max(res.maxLatency, latency)
	res.minLatency = math.Min(res.minLatency, latency)
	res.avgLatency = (res.avgLatency * float64(res.noOfRequestCompleted) + latency) / float64(res.noOfRequestCompleted + 1)
	res.noOfRequestCompleted += 1
	if isSuccess {
		res.noOfSuccessfulRequest += 1
	}
}

func (res *Result) PrintResult() {
	fmt.Printf("maxLatency: %0.2f\n", res.maxLatency)
	fmt.Printf("minLatency: %0.2f\n", res.minLatency)
	fmt.Printf("avgLatency: %0.2f\n", res.avgLatency)
	fmt.Printf("noOfRequestCompleted: %d\n", res.noOfRequestCompleted)
	fmt.Printf("noOfSuccessfulRequest: %d\n", res.noOfSuccessfulRequest)
}

func (res *Result) GetNoOfRequestcompleted() int {
	return res.noOfRequestCompleted
}