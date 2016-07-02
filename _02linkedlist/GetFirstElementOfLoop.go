package main

import (
	ll "../libs/linkedlist"
	"fmt"
)

func main() {
	head := &ll.LinkedList{nil, 1}
	head.Add(2).Add(3).Add(4).Add(5).Add(6).Add(7).Add(8).Add(9)

	node := head.SearchNode(2)
	head.AddNode(node)

	loopHead := GetFirstElemntOfLoop(head)

	if (loopHead != nil) {
		fmt.Println(loopHead.Data)
	}

}

func GetFirstElemntOfLoop(head *ll.LinkedList) *ll.LinkedList {
	slow := head
	fast := head

	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next.Next

		if (slow == fast) {
			break;
		}
	}
	slow = head

	for ; fast != nil && fast.Next != nil; {
		slow = slow.Next
		fast = fast.Next
		if (slow == fast) {
			return slow
		}
	}

	return nil
}




