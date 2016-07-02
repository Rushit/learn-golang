package linkedlist

import (
	"fmt"
	"strconv"
)

type LinkedList struct {
	Next *LinkedList
	Data int
}

func (l *LinkedList) Add(data int) *LinkedList {
	newNode := &LinkedList{nil, data}
	currentNode := l;
	for ; currentNode.Next != nil; {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode

	return l
}

func (l *LinkedList) SearchNode(data int) *LinkedList {
	currentNode := l;
	for ; currentNode != nil; {
		if (currentNode.Data == data) {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil

}

func (l *LinkedList) Delete(node *LinkedList) *LinkedList {
	if (l.Data == node.Data) {
		return l.Next
	}
	currentNode := l;
	nextNode := l.Next
	for ; nextNode != nil; {
		if (nextNode.Data == node.Data) {
			currentNode.Next = nextNode.Next
			break
		} else {
			currentNode = currentNode.Next
			nextNode = nextNode.Next
		}
	}
	return l
}

func (l *LinkedList) Print() {
	currentNode := l
	for ; currentNode != nil; {
		fmt.Print(strconv.Itoa(currentNode.Data))
		if(currentNode.Next != nil) {
			fmt.Print(",")
		}
		currentNode = currentNode.Next
	}
	fmt.Println("")
}


func (l *LinkedList) ReversePrint() {
	var newHead *LinkedList = nil
	currentNode := l
	for ; currentNode != nil; {
		if(newHead == nil) {
			newHead = &LinkedList{nil, currentNode.Data}
		} else {
			newNode := &LinkedList{nil, currentNode.Data}
			newNode.Next = newHead
			newHead = newNode
		}
		currentNode = currentNode.Next
	}
	newHead.Print()
}
