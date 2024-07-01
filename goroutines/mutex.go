package goroutines

import (
	"fmt"
	"sync"
)

type Container struct{
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) increase(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func Mutex() {
	c:= Container{
		counters: map[string]int{
			"a": 0,
			"b": 0,
		},
	}

	var wg sync.WaitGroup
	//closure
	doIncrement := func(name string, n int) {
		for i:= 0; i<n; i++ {
			c.increase(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	wg.Wait()

	fmt.Println(c.counters)
}