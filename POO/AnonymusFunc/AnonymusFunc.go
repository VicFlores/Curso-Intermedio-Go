package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5

	/* func anónima */
	y := func() int {
		return x * 2
	}()

	fmt.Println("Y:", y)

	/*  */
	c := make(chan int)

	go func() {
		fmt.Println("Iniciando function")
		time.Sleep(5 * time.Second)
		fmt.Println("Finalizando function")
		c <- 69
	}()

	<-c
}
