package counter

import (
	"testing"
	"time"
)

func TestNewRequestCounter(t *testing.T) {
	rc := NewRequestCounter()
	if rc == nil {
		t.Error("NewRequestCounter() returned nil")
	}
	if len(rc.RequestTimes) != 0 {
		t.Errorf("Expected empty RequestTimes, got %v", rc.RequestTimes)
	}
}

func TestRecordAndCount(t *testing.T) {
	rc := NewRequestCounter()

	// Record a single request
	count := rc.RecordAndCount()
	if count != 1 {
		t.Errorf("Expected count to be 1, got %d", count)
	}

	// Record multiple requests
	for i := 0; i < 5; i++ {
		rc.RecordAndCount()
	}

	if len(rc.RequestTimes) != 6 {
		t.Errorf("Expected 6 recorded requests, got %d", len(rc.RequestTimes))
	}

	// Test counting with time cutoff (60 seconds)
	time.Sleep(65 * time.Second)
	countAfterSleep := rc.RecordAndCount()
	if countAfterSleep != 1 {
		t.Errorf("Expected count to be 1 after sleep, got %d", countAfterSleep)
	}
}
