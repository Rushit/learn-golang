package main

import "fmt"

func main() {
	cache := make(map[int]int)
	fmt.Println("Top Down : ", fib(10, cache))
	fmt.Println("Bottom Up: ", fib2(10))
}

func fib(n int, cache map[int]int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if val, ok := cache[n]; ok {
		return val
	} else {
		cache[n] = fib(n - 1, cache) + fib(n - 2, cache)
		return cache[n]
	}
}

func fib2(n int) int {
	var table = make([]int, n+1)

	table[0] = 0
	table[1] = 1

	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}

	return table[n]
}
