package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeSortedLists(list1, list2).Display() // Output : *ListNode => 1, 1, 2, 3, 4, 4

	mergeSortedListsWithRecursive(list1, list2).Display() // Output : *ListNode => 1, 1, 2, 3, 4, 4
}

func mergeSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var tail *ListNode

	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil {
			if tail == nil {
				result = list2
				tail = list2
			} else {
				tail.Next = list2
			}
			break
		}

		if list2 == nil {
			if tail == nil {
				result = list1
				tail = list1
			} else {
				tail.Next = list1
			}
			break
		}

		var listVal int

		if list1.Val <= list2.Val {
			listVal = list1.Val
			list1 = list1.Next
		} else {
			listVal = list2.Val
			list2 = list2.Next
		}

		node := &ListNode{
			Val:  listVal,
			Next: nil,
		}

		if result == nil {
			result = node
		} else {
			tail.Next = node
		}
		tail = node
	}

	return result
}

func mergeSortedListsWithRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeSortedListsWithRecursive(list1.Next, list2)
		return list1
	}
	list2.Next = mergeSortedListsWithRecursive(list1, list2.Next)
	return list2
}

func (l *ListNode) Display() {
	list := l
	for list != nil {
		fmt.Printf("%+v -> ", list.Val)
		list = list.Next
	}
	fmt.Println()
}
