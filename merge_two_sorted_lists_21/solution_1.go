package merge_two_sorted_lists_21

// * Definition for singly-linked list. //
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			list1, list2 = list2, list1
		}
		list1.Next = mergeTwoLists(list1.Next, list2)
	}
	if list1 == nil {
		return list2
	}
	return list1
}
