package main

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) Append(value interface{}) *List {
	node := &Node{
		next:  nil,
		value: value,
	}

	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}

	l.tail = node

	return l
}

func (l *List) Prepend(value interface{}) *List {
	node := &Node{
		next:  l.head,
		value: value,
	}

	l.head = node

	return l
}

func (l *List) Insert(index int, value interface{}) *List {
	p := l.head
	i := 0

	for p.next != nil {
		if index == i {
			node := &Node{
				next:  p.next,
				value: p.value,
			}

			p.value = value
			p.next = node

			break
		}

		p = p.next
		i++
	}

	return l
}

func (l *List) Remove(index int) *List {
	p := l.head
	i := 0

	for p.next != nil {
		if index == i {
			next := p.next
			p.value = next.value
			p.next = next.next

			break
		}

		p = p.next
		i++
	}

	return l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v -> ", list.value)
		list = list.next
	}
	fmt.Println()
}

func main() {
	list := List{}
	list.Append(10).Append(20).Append(9).Append(5).Prepend(2).Remove(1).Display() // Output: 2 -> 20 -> 9 -> 5 ->
}
