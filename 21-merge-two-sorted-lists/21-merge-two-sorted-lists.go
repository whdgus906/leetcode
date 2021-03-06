/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 != nil && (l1.Val > l2.Val) {
		l1, l2 = l2, l1

	}
    
	if l1 != nil {
		l1.Next = mergeTwoLists(l1.Next, l2)
	}
    
	return l1
}