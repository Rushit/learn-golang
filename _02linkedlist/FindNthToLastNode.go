package main

import (
	ll "../libs/linkedlist"
	"fmt"
)

func main() {

	l := &ll.LinkedList{nil, 1}
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(8)
	l.Add(9)
	l.Add(10)
	node1 := FindNthToLastNode(l, 5)
	if (node1 != nil) {
		fmt.Println(node1.Data)
	} else {
		fmt.Println("nil")
	}

	node2 := FindNthToLastNode(l, 11)
	if (node2 != nil) {
		fmt.Println(node2.Data)
	} else {
		fmt.Println("nil")
	}

}

func FindNthToLastNode(head *ll.LinkedList, n int) *ll.LinkedList {
	ele := head
	last := head

	for i := 0; i < n; i++ {
		if (last == nil) {
			return nil
		}
		last = last.Next
	}

	for ; last != nil; {
		last = last.Next
		ele = ele.Next
	}

	return ele
}
