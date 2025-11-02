package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
)

type Counter struct {
	count int
}

func (a *Counter) Increment() {
	mu.Lock()
	a.count += 1
	defer mu.Unlock()
}

func (a *Counter) Value() int {
	mu.Lock()
	defer mu.Unlock()
	return a.count
}

func main() {
	wg.Add(100)
	var counter Counter
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终计数:", counter.Value())
}
