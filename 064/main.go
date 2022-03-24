// write a program that
// ○ launches 10 goroutines
// ■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them
package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					c <- i*10 + m
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
