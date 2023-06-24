package main

import (
	"log"
	"net/http"
	"sync"
)

func go_get(link string, wg *sync.WaitGroup, mut *sync.Mutex, result *[]int) {
	// tell wait group that we just finished one go routine, thus decrementing count
	defer wg.Done()

	resp, err := http.Get(link)
	if err != nil {
		log.Println("got error from: ", link, "error: ", err)
	}
	log.Println("success mesage from: ", link, resp.StatusCode)

	// lock variable/memory before performing read/write operation,
	// to prevent from deadlock/race condition
	// and unlock it after operation
	mut.Lock()
	*result = append(*result, resp.StatusCode)
	mut.Unlock()
	
	// log.Println("print: ",len(*result))
}

func main() {
	log.Println("start")
	var wg sync.WaitGroup // to keep count of routines
	var mut sync.Mutex // for lock to memory

	var results []int
	links := []string{
		"https://www.google.com/",
		"https://www.fb.com/",
		"https://www.github.com/",
	}
	for i := 0; i < len(links); i++ {
		// it will create go routines(thread)
		go go_get(links[i], &wg, &mut, &results)
		// it will keep count of go routines, thus incrementing count
		wg.Add(1) 
	}

	// wait until all routines get finished i.e. routine num=0
	wg.Wait()

	log.Println("result len: ",len(results))
}
