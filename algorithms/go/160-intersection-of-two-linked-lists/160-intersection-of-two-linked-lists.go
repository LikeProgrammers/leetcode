package getIntersectionNode

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ml := map[*ListNode]bool{}
	for p := headA; p != nil; p = p.Next {
		ml[p] = true
	}
	for p := headB; p != nil; p = p.Next {
		if _, found := ml[p]; found {
			return p
		}
	}
	return nil
}
