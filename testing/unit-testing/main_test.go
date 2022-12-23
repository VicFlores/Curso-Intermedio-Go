package main

import "testing"

func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.n)
		}
	}

}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{3, 2, 3},
		{5, 4, 5},
		{3, 6, 6},
		{7, 9, 9},
	}

	for _, item := range tables {
		max := GetMax(item.a, item.b)

		if max != item.n {
			t.Errorf("Get max was incorrect, got %d expected %d", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		fib := Fibonacci(item.a)

		if fib != item.n {
			t.Errorf("Fibonacci incorrect, got %d expected %d", fib, item.n)
		}
	}
}
