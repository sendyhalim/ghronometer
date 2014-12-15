package timer

import (
	"testing"
	"time"
)

func assert(t *testing.T, first interface{}, second interface{}) {
	if first != second {
		t.Errorf("%v != %v", first, second)
	}
}

func createTimer() *Timer {
	return &Timer{}
}

func TestTimeout(t *testing.T) {
	timer := createTimer()
	timer.Timeout = 40
	timer.Start()
	time.Sleep(1e8)
	assert(t, 40, timer.Counter())

	_, open := <-timer.counterChannel
	assert(t, open, false)
}

func TestStop(t *testing.T) {
	timer := createTimer()
	timer.Start()
	timer.Stop()
	_, open := <-timer.counterChannel
	assert(t, open, false)
}
