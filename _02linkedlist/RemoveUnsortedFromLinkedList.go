package main

import (
	ll "../libs/linkedlist"
)

func main() {
	l := &ll.LinkedList{nil, 1}
	l.Add(2)
	l.Add(4)
	l.Add(1)
	l.Add(4)
	l.Add(3)
	l.Add(2)
	l.Add(1)
	l.Add(1)
	l.Print()
	RemoveDups(l)
	l.Print()

}

func RemoveDups(head *ll.LinkedList) {
	dataMap := make(map[int]bool)
	var prevNode *ll.LinkedList = nil
	currentNode := head
	for ; currentNode != nil; {
		if (hasDups(currentNode, dataMap)) {
			prevNode.Next = currentNode.Next
		} else {
			dataMap[currentNode.Data] = true
			prevNode = currentNode
		}
		currentNode = currentNode.Next
	}
}

func hasDups(node *ll.LinkedList, nodeMap map[int]bool) bool {
	return nodeMap[node.Data] == true
}


