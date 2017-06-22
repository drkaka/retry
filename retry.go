package retry

import (
	"time"
)

// Call to use retry functions
//
// times to retry the function
// interval the initial seconds used to wait for another call
// factor to be mutiplied to interval after every retry
func Call(f func() bool, times, interval int, factor float32) bool {
	if f() {
		return true
	}

	wait := interval
	for index := 0; index < times; index++ {
		time.Sleep(time.Duration(wait) * time.Second)
		wait = int(float32(wait) * factor)

		if f() {
			return true
		}
	}
	return false
}
