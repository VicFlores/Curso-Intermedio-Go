package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Id: %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("Id: %d finished\n", i)

	<-c
}

func main() {
	/* Buffered channels */
	var wg sync.WaitGroup
	c := make(chan int, 5)

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)

		go doSomething(i, &wg, c)
	}

	wg.Wait()

}
