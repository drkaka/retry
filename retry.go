package retry

import (
	"time"
)

// Call to use retry functions
//
// f (times int) bool will passin current retry times and if success, return true
// times to retry the function
// interval the initial seconds used to wait for another call
// factor to be mutiplied to interval after every retry
func Call(f func(int) bool, times, interval int, factor float32) bool {
	if f(0) {
		return true
	}

	wait := interval
	for i := 0; i < times; i++ {
		time.Sleep(time.Duration(wait) * time.Second)
		wait = int(float32(wait) * factor)

		if f(i + 1) {
			return true
		}
	}
	return false
}
