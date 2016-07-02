package main

import (
	ll "../libs/linkedlist"
)

func main() {
	l := &ll.LinkedList{nil, 1}
	l.Add(12).Add(31).Add(4).Add(23).Add(10).Add(5).Add(51)
	l.Print()
	l = PartitionLinkedList(l, 23)
	l.Print()
}

func PartitionLinkedList(head *ll.LinkedList, data int) *ll.LinkedList {
	var leftStart  *ll.LinkedList = nil
	var rightStart *ll.LinkedList = nil
	var middleNode *ll.LinkedList = nil

	currentNode := head
	for ; currentNode != nil; {
		next := currentNode.Next
		if (currentNode.Data == data) {
			middleNode = currentNode
		} else {
			if (currentNode.Data < data) {
				currentNode.Next = leftStart
				leftStart = currentNode
			} else {
				currentNode.Next = rightStart
				rightStart = currentNode
			}
		}
		currentNode = next
	}

	if(middleNode != nil){
		middleNode.Next = rightStart
		rightStart = middleNode
	}

	if (leftStart == nil) {
		return rightStart
	}

	leftEnd := leftStart
	for ; leftEnd.Next != nil; {
		leftEnd = leftEnd.Next
	}
	leftEnd.Next = rightStart

	return leftStart

}
