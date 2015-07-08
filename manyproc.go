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

			log.Printf("now receiving start %v, from chan %v", i, c)
			select {
			case cmd := <-c:
				log.Printf("OK: command received: %v", cmd)
			case <-time.After(100 * time.Second):
				log.Printf("FAIL: %v is timeout", i)
			}
			return
		}(i, ch)
	}
	log.Printf("now sending start")
	for i, ch := range chans {
		log.Printf("sending %v, %v",i, ch)
		ch <- i
	}
	wg.Wait()
	log.Printf("now all processes completed")
}
