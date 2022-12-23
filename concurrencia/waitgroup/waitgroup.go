package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Inicio proceso %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizo proceso")
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}
