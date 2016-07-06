package main

import (
	"fmt"
	"strconv"
)

func main() {
	var lineStr string
	fmt.Scanln(&lineStr)
	lines, _ := strconv.Atoi(lineStr)
	array := []int{}
	for i := 0; i < lines; i++ {
		var numStr string
		fmt.Scanln(&numStr)
		num, _ := strconv.Atoi(numStr)
		array = append(array, num)
	}
	fmt.Println(LIS_NLGN(array))
}

func LIS_N2(array []int) int {

	lis := make([]int, len(array))
	for aLis := range lis {
		lis[aLis] = 1
	}
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[i] > array[j] && lis[i] < lis[j] + 1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	return getMax(lis)
}

func LIS_NLGN(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return 1
	}
	size := len(a)
	tt := make([]int, len(a))
	len := 1
	tt[0] = a[0]
	for i := 0; i < size; i++ {
		if a[i] < tt[0] {
			tt[0] = a[i]
		} else if a[i] > tt[len - 1] {
			tt[len] = a[i]
			len++
		} else {
			bsIdx := binarySearchIndex(tt, -1, len - 1, a[i])
			tt[bsIdx] = a[i]
		}
	}
	return len
}

func binarySearchIndex(tt []int, l int, r int, key int) int {
	for ; r - l > 1; {
		m := l + (r - l) / 2
		if tt[m] >= key {
			r = m
		} else {
			l = m
		}

	}
	return r
}

func getMax(lis []int) int {
	if len(lis) == 0 {
		return 0
	}
	max := lis[0]
	for i := range lis {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

