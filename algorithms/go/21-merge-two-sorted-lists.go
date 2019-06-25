package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{0, nil}
	pl1 := l1
	l1_nodes := []int{1, 2, 4}
	for _, v := range l1_nodes {
		node := ListNode{v, nil}
		pl1.Next = &node
		pl1 = pl1.Next
	}
	l1 = l1.Next

	l2 := &ListNode{0, nil}
	pl2 := l2
	l2_nodes := []int{1, 3, 4}
	for _, v := range l2_nodes {
		node := ListNode{v, nil}
		pl2.Next = &node
		pl2 = pl2.Next
	}
	l2 = l2.Next

	// printListNode(head)

	ret := mergeTwoLists(l1, l2)
	printListNode(ret)
}

func printListNode(head *ListNode) {
	prun := head
	for ; prun != nil; prun = prun.Next {
		fmt.Println(prun.Val)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pl1 := l1
	pl2 := l2
	l := &ListNode{0, nil}
	pl := l
	for ; pl1 != nil || pl2 != nil; pl = pl.Next {
		if pl1 != nil && pl2 == nil {
			node := ListNode{pl1.Val, nil}
			pl.Next = &node
			pl1 = pl1.Next
		} else if pl2 != nil && pl1 == nil {
			node := ListNode{pl2.Val, nil}
			pl.Next = &node
			pl2 = pl2.Next
		} else {
			if pl1.Val <= pl2.Val {
				node := ListNode{pl1.Val, nil}
				pl.Next = &node
				pl1 = pl1.Next
			} else {
				node := ListNode{pl2.Val, nil}
				pl.Next = &node
				pl2 = pl2.Next
			}
		}
	}

	return l.Next
}
