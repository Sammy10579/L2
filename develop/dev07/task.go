package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	defer close(out)

	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, c := range channels {
		go func(c <-chan interface{}) {
			for val := range c {
				out <- val
			}
			wg.Done()
		}(c)
	}

	wg.Wait()

	return out

}

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
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
