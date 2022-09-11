package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type Stack struct {
	top    *Node
	length int
}

func (s *Stack) Push(value interface{}) *Stack {
	node := &Node{
		next:  s.top,
		value: value,
	}

	s.top = node
	s.length++

	return s
}

func (s *Stack) Pop() *Stack {
	s.top.value = s.top.next.value
	s.top.next = s.top.next.next

	s.length--

	return s
}

func (s *Stack) Display() {
	top := s.top

	for top != nil {
		fmt.Printf("%+v => ", top.value)
		top = top.next
	}

	fmt.Println()
}

func (s *Stack) Length() {
	fmt.Printf("List length is : %d \n", s.length)
}

func (s *Stack) Peek() {
	fmt.Printf("Peek is : %+v \n", s.top.value)
}

func main() {
	stack := Stack{}
	s := stack.Push("Google").Push("Facebook").Push("Youtube").Push("Twitter").Pop()

	s.Length()  // Output: List length is : 3
	s.Display() // Output: Youtube => Facebook => Google =>
	s.Peek()    // Output: Peek is : Youtube
}
