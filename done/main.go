package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(4*time.Second),
		sig(3*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))

}

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	done := make(chan struct{})

	for _, ch := range channels {
		go func(c <-chan interface{}) {
			select {
			case <-c:
				select {
				case <-done:
				default:
					close(done)
					close(out)
				}
			case <-done:
			}
		}(ch)
	}

	return out
}
