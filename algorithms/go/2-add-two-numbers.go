package main

import "fmt"

func main() {
    // Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    // Output: 7 -> 0 -> 8
    // Explanation: 342 + 465 = 807.

    list1 := []int{2, 4, 3}
    list2 := []int{5, 6, 7}
    head1 := ListNode{0, nil}
    head2 := ListNode{0, nil}
    pcurrent1 := &head1
    pcurrent2 := &head2
    for i := 0; i < len(list1); i++ {
        pcurrent1.Next = &ListNode{list1[i], nil}
        pcurrent1 = pcurrent1.Next
    }
    for i := 0; i < len(list2); i++ {
        pcurrent2.Next = &ListNode{list2[i], nil}
        pcurrent2 = pcurrent2.Next
    }

    l1 := head1.Next
    l2 := head2.Next
    ret := addTwoNumbers(l1, l2)

    for pcurrent := ret; pcurrent != nil; pcurrent = pcurrent.Next {
        fmt.Println(pcurrent.Val)
    }
}

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := ListNode{0, nil}
    pcurrent := &head

    carry := 0

    pl1 := l1
    pl2 := l2

    for pl1 != nil || pl2 != nil || carry != 0 {

        l1_val := 0
        l2_val := 0
        if pl1 != nil {
            l1_val = pl1.Val
            pl1 = pl1.Next
        }
        if pl2 != nil {
            l2_val = pl2.Val
            pl2 = pl2.Next
        }

        sum := l1_val + l2_val + carry
        current := sum % 10
        carry = sum / 10
        pcurrent.Next = &ListNode{current, nil}
        pcurrent = pcurrent.Next
    }

    return head.Next
}
