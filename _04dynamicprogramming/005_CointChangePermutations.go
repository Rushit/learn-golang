package main

import (
	"os"
	"fmt"
	"bufio"
)

/**
	Problem: https://www.hackerrank.com/challenges/coin-change
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)
	var c []int = make([]int, m)
	for i := range c {
		fmt.Fscan(reader, &c[i])
	}

	fmt.Println(CoinChange(n, c))
}

func CoinChange(total int, coins []int) int {
	temp := make([]int, total + 1)

	temp[0] = 1;
	for i := 0; i < len(coins); i++ {
		for value := 0; value <= total; value++ {
			if value >= coins[i] {
				idx := value - coins[i]
				temp[value] += temp[idx]
			}
			fmt.Println(coins[i], value, temp)
		}
		fmt.Println("")
	}
	return temp[total];
}
