package main

import "fmt"

/** PROBLEM STATEMENT:
A child is running up a staircase with n steps, and can hop either 1 step, 2 steps, or
3 steps at a time. Implement a method to count how many possible ways the child
can run up the stairs.
 */

func main() {
	countMap := make(map[int]int)
	fmt.Print("Total ways are -> ")
	fmt.Println(CountSteps(18, countMap))
}

func CountSteps(totalSteps int, countMap map[int]int) int {
	if totalSteps < -1 {
		return 0
	} else if totalSteps == 0 {
		return 1
	} else if val, ok := countMap[totalSteps]; ok && val > -1 {
		return val
	} else {
		countMap[totalSteps] = CountSteps(totalSteps - 1, countMap) + CountSteps(totalSteps - 2, countMap) + CountSteps(totalSteps - 3, countMap)
		return countMap[totalSteps]
	}
}
