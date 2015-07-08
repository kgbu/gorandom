package main

import (
	"log"
	"time"
)

func main() {
	var chans [200000]chan int
	for i := range chans {
   		chans[i] = make(chan int)
	}
	for i, ch := range chans {
		go func(i int, c chan int) {
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
}
