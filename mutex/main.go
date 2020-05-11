package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	lock    sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("final counter: %d", counter))
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	counter++
	wg.Done()
}