package main

import (
	"bufio"
	"os"
	"fmt"
)

/**
	Problem: https://www.hackerrank.com/challenges/array-left-rotation
 */
func main() {
	reader := bufio.NewReader(os.Stdin)
	var len int
	var rotate int
	fmt.Fscan(reader, &len)
	fmt.Fscan(reader, &rotate)
	a := make([]int64, len)
	for j := 0; j < len; j++ {
		fmt.Fscan(reader, &a[j])
	}

	a = LeftRotateArray(rotate, a)
	for i := range a {
		fmt.Print(a[i], " ")
	}
}

func LeftRotateArray(rotate int, a []int64) []int64 {
	temp := make([]int64, rotate)
	for i := 0; i < rotate; i++ {
		temp[i] = a[i]
	}
	rotateI := len(a) - rotate
	for i := 0; i < len(a); i++ {
		if i < rotateI {
			a[i] = a[i + rotate]
		} else {
			a[i] = temp[i - rotateI]
		}
	}
	return a
}
