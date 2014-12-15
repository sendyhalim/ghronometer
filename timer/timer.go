package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Timeout   int
	PrintTime bool

	counter        int
	counterChannel chan int
}

func (t *Timer) Start() {
	// in milisecond
	ticker := time.Tick(1e6)
	// restart everything
	t.Reset()

	go func() {
		for {
			// stop if timeout is set and counter has exceed the timeout
			if t.Timeout > 0 && t.counter >= t.Timeout {
				t.Stop()
				break
			}

			<-ticker
			t.counterChannel <- t.counter
			t.counter++
		}
	}()

	go func() {
		for {
			totalTime, open := <-t.counterChannel

			if t.PrintTime {
				if !open {
					fmt.Print("\n")
					break
				}

				fmt.Printf("%d \r", totalTime)
			}
		}
	}()
}

func (t *Timer) Counter() int {
	return t.counter
}

func (t *Timer) Reset() {
	if t.counterChannel != nil {
		close(t.counterChannel)
	}

	t.counterChannel = make(chan int, t.Timeout)
	t.counter = 0
}

func (t *Timer) Stop() {
	close(t.counterChannel)
}
