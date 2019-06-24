package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{0, nil}
	prun := head
	nodes := []int{1, 2, 3, 4, 5}
	for _, v := range nodes {
		node := ListNode{v, nil}
		prun.Next = &node
		prun = prun.Next
	}
	head = head.Next

	// printListNode(head)

	n := 6
	ret := removeNthFromEnd(head, n)
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}

	i := 1
	prun := head
	pstep := head

	for prun.Next != nil {
		i++
		prun = prun.Next
		if i > n+1 {
			pstep = pstep.Next
		}
	}

	if n > i {
		return head
	} else if i == n {
		return head.Next
	} else {
		pstep.Next = pstep.Next.Next
		return head
	}
}
