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
}
