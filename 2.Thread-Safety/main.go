package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
	Inc   func()
	Value func() int
}

func NewSafeCounter() *SafeCounter {
	sc := &SafeCounter{}

	sc.Inc = func() {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		sc.count++
	}

	sc.Value = func() int {
		sc.mu.Lock()
		defer sc.mu.Unlock()
		return sc.count
	}
	return sc
}

func main() {
	counter := NewSafeCounter()

	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("Final count:", counter.Value())
}
