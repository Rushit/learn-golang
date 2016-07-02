package main

import (
	ll "../libs/linkedlist"
	"strconv"
	"strings"
)

/**
  6,1,7
+ 3,9,5
--------
1,0,1,2
 */

func main() {
	num1 := &ll.LinkedList{nil, 7}
	num1.Add(1).Add(6)

	num2 := &ll.LinkedList{nil, 5}
	num2.Add(9).Add(3)
	num1.ReversePrint()
	num2.ReversePrint()
	sum := SumNumbersWithLinkedList(num1, num2)
	sum.ReversePrint()
}

func SumNumbersWithLinkedList(num1 *ll.LinkedList, num2 *ll.LinkedList) *ll.LinkedList {
	var sumLL *ll.LinkedList = nil
	var sumHead *ll.LinkedList = nil

	numItr1 := num1
	numItr2 := num2
	var carry int = 0
	for ; numItr1 != nil || numItr2 != nil || carry > 0; {

		sum := sum(numItr1, numItr2, carry)

		sumDigit := sum
		carry = 0
		if (sum > 9) {
			sumStr := strconv.Itoa(sum) //guaranted to be two digit only
			s := strings.Split(sumStr, "")
			carry, _ = strconv.Atoi(s[0])
			sumDigit, _ = strconv.Atoi(s[1])
		}

		if (sumLL == nil) {
			sumLL = &ll.LinkedList{nil, sumDigit}
			sumHead = sumLL
		} else {
			newNode := &ll.LinkedList{nil, sumDigit}
			sumLL.Next = newNode
			sumLL = sumLL.Next
		}

		if (numItr1 != nil) {
			numItr1 = numItr1.Next
		}
		if (numItr2 != nil) {
			numItr2 = numItr2.Next
		}
	}
	return sumHead
}

func sum(num1 *ll.LinkedList, num2 *ll.LinkedList, carry int) int {
	sum := 0

	if (num1 != nil) {
		sum += num1.Data
	}

	if (num2 != nil) {
		sum += num2.Data
	}

	return sum + carry
}