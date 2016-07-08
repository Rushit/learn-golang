package main

import (
	"bufio"
	"os"
	"fmt"
)

/**
	Problem: https://www.hackerrank.com/challenges/stockmax
	Formula:   max   {a[i] - a[j]} , for all i
			  j > i

 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	var testCaseCount int
	fmt.Fscan(reader, &testCaseCount)
	allArrays := [][]int64{}
	for i := 0; i < testCaseCount; i++ {
		var num int
		fmt.Fscan(reader, &num)
		aArray := []int64{}
		for j := 0; j < num; j++ {
			var n int64
			fmt.Fscan(reader, &n)
			aArray = append(aArray, n)
		}
		allArrays = append(allArrays, aArray)
	}
	for i := range allArrays {
		fmt.Println(MaximizeStockProfit(allArrays[i]))
	}
}

func MaximizeStockProfit(stockPrice []int64) int64 {

	maxArray := make([]int64, len(stockPrice))
	size := len(stockPrice)
	maxArray[size - 1] = stockPrice[size - 1]

	for i := size - 2; i >= 0; i-- {
		maxArray[i] = max(stockPrice[i], maxArray[i + 1])
	}

	var maxProfit int64 = 0
	for i := range maxArray {
		maxProfit = maxProfit + ( maxArray[i] - stockPrice[i])
	}
	return maxProfit
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
