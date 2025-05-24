package service

import (
	"testing"
	"time"
)

func TestTokenBucket_CheckIsAllowed_WithinCapacity(t *testing.T) {
	bucket := NewTokenBucket(5, 1, time.Second)

	// Should allow first 3 requests
	for i := 0; i < 5; i++ {
		if !bucket.CheckIsAllowed() {
			t.Errorf("Expected request %d to be allowed", i)
		}
	}

	// 4th request should be rate limited
	if bucket.CheckIsAllowed() {
		t.Error("Expected request to be rate-limited")
	}

}

func TestTokenBucket_CheckIsAllowed_Refill(t *testing.T) {
	bucket := NewTokenBucket(2, 1, time.Second)

	bucket.CheckIsAllowed()
	bucket.CheckIsAllowed()

	time.Sleep(time.Second * 1)

	if !bucket.CheckIsAllowed() {
		t.Error("Expected token to be refilled and request should be allowed")
	}

}

func TestTokenBucket_ConcurrentAccess(t *testing.T) {
	bucket := NewTokenBucket(10, 10, time.Second)

	success := 0
	done := make(chan bool)

	for i := 0; i < 20; i++ {
		go func() {
			if bucket.CheckIsAllowed() {
				success++
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for range 20 {
		<-done
	}

	if success > 10 {
		t.Errorf("Expected at most 10 allowed requests, got %d", success)
	}
}
