package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	"math/big"
)

/**
	Problem Def: https://www.hackerrank.com/challenges/fibonacci-modified
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	lineStr, _ := reader.ReadString('\n')
	lineStr = strings.TrimRight(lineStr, "\n")
	nums := strings.Split(lineStr, " ")
	a, _ := strconv.Atoi(nums[0])
	b, _ := strconv.Atoi(nums[1])
	n, _ := strconv.Atoi(nums[2])

	ans := FibonacciModified(a, b, n)
	fmt.Println(ans.String())

}

func FibonacciModified(a int, b int, n int) big.Int {
	arr := make([]big.Int, n, n)
	arr[0] = *big.NewInt(int64(a))
	arr[1] = *big.NewInt(int64(b))

	for i := 2; i < n; i++ {
		var i1 *big.Int = &arr[i - 1]
		var i2 *big.Int = &arr[i - 2]
		m := big.NewInt(0).Mul(i1, &arr[i - 1])
		a := big.NewInt(0).Add(m, i2)
		arr[i] = *a
	}

	return arr[n - 1]
}
