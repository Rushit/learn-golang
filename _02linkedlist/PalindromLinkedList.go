package main

import (
	ll "../libs/linkedlist"
	"fmt"
)

func main() {
	head := &ll.LinkedList{nil, 1}
	head.Add(2).Add(3).Add(4).Add(3).Add(2).Add(1)
	fmt.Println(IsPalindromeLinkedList(head))
}

func IsPalindromeLinkedList(list *ll.LinkedList) bool {
	var reverseList *ll.LinkedList = nil

	currentNode := list
	for ; currentNode != nil;  {
		if(reverseList == nil) {
			reverseList = &ll.LinkedList{nil, currentNode.Data}
		} else {
			newNode := &ll.LinkedList{reverseList, currentNode.Data }
			reverseList = newNode
		}
		currentNode = currentNode.Next
	}

	currentNode = list
	for ; currentNode != nil;  {
		if(reverseList.Data != currentNode.Data) {
			return false
		}
		reverseList = reverseList.Next
		currentNode = currentNode.Next
	}

	return true
}


