package counter

import "sync"

type RequestCounter struct {
	sync.Mutex
	l []int64
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{}
}
