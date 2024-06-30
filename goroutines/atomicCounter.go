package goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounter() {
	fmt.Println("in go routine");
	var counter atomic.Uint32	//1st method introduce a atomic variable
	var mu sync.Mutex	//using mutex in a normal var
	var count int32

	var wg sync.WaitGroup

	for i:= 0 ; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c:= 0; c< 1000; c++ {
				counter.Add(1)
				mu.Lock()
				count++
				mu.Unlock();
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter count: ", counter.Load())
	fmt.Println("count count: ", count)
}