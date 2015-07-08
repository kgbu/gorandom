package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var chans [200000]chan int
	var wg sync.WaitGroup

	for i := range chans {
   		chans[i] = make(chan int)
	}
	for i, ch := range chans {
		wg.Add(1)
		go func(i int, c chan int) {
			defer wg.Done()

			select {
			case cmd := <-c:
			case <-time.After(100 * time.Second):
				log.Printf("FAIL: %v is timeout", i)
			}
			return
		}(i, ch)
	}
	for i, ch := range chans {
		ch <- i
	}
	wg.Wait()
}
