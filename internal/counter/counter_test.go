package counter

import (
	"testing"
	"time"
)

func TestRequestCounter_RecordAndCount(t *testing.T) {
	c := NewRequestCounter()

	// Edge Case 1: No requests
	if count := c.RecordAndCount(); count != 0 {
		t.Errorf("Expected 0 requests, got %d", count)
	}

	// Adding a request exactly 60 seconds ago
	pastTime := time.Now().Add(-60 * time.Second).Unix()
	c.RequestTimes = append(c.RequestTimes, pastTime)
	if count := c.RecordAndCount(); count != 1 {
		t.Errorf("Expected 1 request, got %d", count)
	}

	// Adding a request just outside the 60-second window
	c.RequestTimes = append(c.RequestTimes, pastTime-1)
	if count := c.RecordAndCount(); count != 1 {
		t.Errorf("Expected 1 request (ignoring the one just outside 60 seconds), got %d", count)
	}

	// Clear the counter for the next test
	c = NewRequestCounter()

	// Adding multiple requests
	for i := 0; i < 5; i++ {
		c.RequestTimes = append(c.RequestTimes, time.Now().Unix()-int64(i*10))
	}
	if count := c.RecordAndCount(); count != 5 {
		t.Errorf("Expected 5 requests, got %d", count)
	}
}
