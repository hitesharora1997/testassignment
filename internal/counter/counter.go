package counter

import (
	"fmt"
	"sync"
	"time"
)

type RequestCounter struct {
	sync.Mutex
	RequestTimes []int64
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{}
}

func (request *RequestCounter) RecordAndCount() int {
	request.Lock()
	defer request.Unlock()

	now := time.Now().Unix()

	request.RequestTimes = append(request.RequestTimes, now)

	cutoff := now - 60
	index := 0

	for i, t := range request.RequestTimes {
		fmt.Println("t", t)
		if t > cutoff {
			index = i
			break
		}
	}
	request.RequestTimes = request.RequestTimes[index:]

	return len(request.RequestTimes)
}
