// Fix the race condition you created in the previous exercise by using a mutex
// ● it makes sense to remove runtime.Gosched()
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
