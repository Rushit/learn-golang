package main
/**
	Problem: https://www.hackerrank.com/challenges/maxsubarray
 */
import (
	"bufio"
	"os"
	"fmt"
)

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
	for aRow := range allArrays {
		fmt.Println(MaximumSubArraySum(allArrays[aRow]))
	}
}

func MaximumSubArraySum(a []int64) (int64, int64) {
	var maxSumNonCont int64 = a[0]
	var maxSumC int64 = a[0]
	var tempSumC int64 = a[0]
	for i := range a {
		if i == 0 {
			continue
		}

		maxSumNonCont = max(maxSumNonCont + a[i], max(maxSumNonCont, a[i]))

		if tempSumC < 0 {
			tempSumC = 0
		}
		tempSumC = tempSumC + a[i]
		maxSumC = max(maxSumC, tempSumC)

	}

	return maxSumC, maxSumNonCont
}

func max(a int64, b int64) int64 {
	if a < b {
		return b
	} else {
		return a
	}
}
