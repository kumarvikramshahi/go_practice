// program to just increase count using channels.
// increse cnt through one routine and display through other routine

package main

import (
	"log"
	"sync"
)

func send(wg *sync.WaitGroup, val chan<- int) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		// sending values through channel to other go routines
		val <- i
	}
	// closing channel to tell other routines to not wait for values
	close(val)
}

func receive(wg *sync.WaitGroup, val <-chan int) {
	defer wg.Done()
	for {
		// receiving values through channel 
		res, isopen := <-val
		// exit loop when channel is closed
		if !isopen {
			log.Println("channel closed")
			return
		}
		log.Println("val: ", res)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	val := make(chan int)

	wg.Add(2)
	go send(wg, val)
	go receive(wg, val)
	wg.Wait()
}
