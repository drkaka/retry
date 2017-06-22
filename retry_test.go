package retry

import (
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	start := time.Now().Unix()
	ok := Call(needRetry, 4, 1, 2.0)

	if ok {
		t.Error("should not ok")
	}
	duration := time.Now().Unix() - start

	if duration != int64(15) {
		t.Error("duration wrong")
	}
}

func needRetry() bool {
	return false
}
