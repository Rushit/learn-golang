package main

import "fmt"

func main() {
	array := []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	fmt.Println(LIS(array))
}

func LIS(array []int) int {

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

func getMax(lis []int) int {
	max := lis[0]
	for i := range lis {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

