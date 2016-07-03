package queue

import (
	ll "../../libs/linkedlist"
)

type Queue struct {
	front *ll.LinkedList
	rear *ll.LinkedList
	Size int
}

func (q *Queue) Enqueue(data int)  {
	newNode := &ll.LinkedList{nil, data}
	if(q.front == nil) {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
	q.Size++
}

func (q *Queue) Dequeue() (int, error)  {
	if(q.front == nil) {
		return nil, error{"Queue is empty"}
	}
	result := q.front.Data
	q.front = q.front.Next
	q.Size--
	return result
}

