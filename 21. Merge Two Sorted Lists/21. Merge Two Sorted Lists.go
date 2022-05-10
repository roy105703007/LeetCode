//my answer
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := new(ListNode)
	head := pre
	for list1 != nil || list2 != nil {
		if list1 == nil {
			pre.Next = list2
			pre = pre.Next
			list2 = list2.Next
			continue
		} else if list2 == nil {
			pre.Next = list1
			pre = pre.Next
			list1 = list1.Next
			continue
		}
		if list2.Val > list1.Val {
			pre.Next = list1
			pre = pre.Next
			list1 = list1.Next
		} else {
			pre.Next = list2
			pre = pre.Next
			list2 = list2.Next
		}
	}
	return head.Next
}

//very good answer
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}