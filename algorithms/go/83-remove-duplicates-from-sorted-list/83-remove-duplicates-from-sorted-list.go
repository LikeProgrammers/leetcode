package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{0, nil}
	pl1 := l1
	l1_nodes := []int{1, 1, 2, 3, 3}
	for _, v := range l1_nodes {
		node := ListNode{v, nil}
		pl1.Next = &node
		pl1 = pl1.Next
	}
	head := l1.Next

	// printListNode(head)

	ret := deleteDuplicates(head)
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
func deleteDuplicates(head *ListNode) *ListNode {
	phead := head
	rl := &ListNode{0, nil}
	pr := rl
	for phead != nil {
		current := phead.Val
		pr.Next = &ListNode{current, nil}
		pr = pr.Next
		phead = phead.Next
		for phead != nil && current == phead.Val {
			phead = phead.Next
		}
	}

	return rl.Next
}
