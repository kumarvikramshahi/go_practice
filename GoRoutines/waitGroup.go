package main

import (
	"log"
	"net/http"
	"sync"
)

func go_get(link string, wg *sync.WaitGroup, mut *sync.Mutex, result *[]int) {
	defer wg.Done()

	resp, err := http.Get(link)
	if err != nil {
		log.Println("got error from: ", link, "error: ", err)
	}
	log.Println("success mesage from: ", link, resp.StatusCode)
	mut.Lock()
	*result = append(*result, resp.StatusCode)
	// mut.Unlock()
	log.Println("print: ",len(*result))
}

func main() {
	log.Println("start")
	var wg sync.WaitGroup
	var mut sync.Mutex

	var results []int
	links := []string{
		"https://www.google.com/",
		"https://www.fb.com/",
		"https://www.github.com/",
	}
	for i := 0; i < len(links); i++ {
		go go_get(links[i], &wg, &mut, &results)
		wg.Add(1)
	}

	wg.Wait()
	log.Println("result len: ",len(results))
}
