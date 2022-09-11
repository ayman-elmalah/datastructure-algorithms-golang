package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type Queue struct {
	top    *Node
	bottom *Node
	length int
}

func (q *Queue) Enqueue(value interface{}) *Queue {
	node := &Node{
		next:  nil,
		value: value,
	}

	if q.length == 0 {
		q.top = node
		q.bottom = node
	}

	q.bottom.next = node
	q.bottom = node

	q.length++

	return q
}

func (q *Queue) Dequeue() *Queue {
	q.top = q.top.next
	q.length--

	return q
}

func (q *Queue) Length() {
	fmt.Printf("Queue length is : %d \n", q.length)
}

func (q *Queue) Peek() {
	fmt.Printf("Queue peek is : %+v \n", q.top.value)
}

func (q *Queue) Display() {
	top := q.top

	for top != nil {
		fmt.Printf("%+v => ", top.value)
		top = top.next
	}

	fmt.Println()
}

func main() {
	q := Queue{}
	q.Enqueue("Google").Enqueue("Facebook").Enqueue("Youtube").Enqueue("Twitter").Dequeue()

	q.Length()  // Output: Queue length is : 3
	q.Display() // Output: Facebook => Youtube => Twitter =>
	q.Peek()    // Output: Queue peek is : Facebook
}
