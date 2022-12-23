package main

import "fmt"

func Suma(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

func PrintNames(values ...string) {

	for _, name := range values {
		fmt.Println("Name:", name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func main() {

	/* Func variadica */

	fmt.Println("Suma con 1 numero", Suma(3))
	fmt.Println("Suma con 2 numero", Suma(4, 4))
	fmt.Println("Suma con 3 numero", Suma(5, 5, 5))
	fmt.Println("Suma con 4 numero", Suma(5, 5, 2, 2))

	PrintNames("Vic")
	PrintNames("Vic Ferman")
	PrintNames("Vic Ferman Flores")
	PrintNames("Vic Ferman Escobar")

	/* Retornos con nombre */

	fmt.Println(getValues(2))

}
