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
