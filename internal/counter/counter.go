package counter

import (
	"fmt"
	"sync"
	"time"
)

type RequestCounter struct {
	sync.Mutex
	l []int64
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{}
}

func (request *RequestCounter) RecordCounter() error {
	request.Lock()
	defer request.Unlock()

	now := time.Now().Unix()
	fmt.Println("time", now)

	return nil
}
