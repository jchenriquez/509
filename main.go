package main

import "fmt"

func fib(N int) int {

	fib1 := 0
	fib2 := 1

	if N == 0 {
		return fib1
	}

	if N == 1 {
		return fib2
	}

	currFib := 1

	for currFib < N {
		tFib := fib2
		fib2 = fib1+fib2
		fib1 = tFib
		currFib++
	}

	return fib2
}

func main() {
	fmt.Printf("N fib: %d\n", fib(20))
}
