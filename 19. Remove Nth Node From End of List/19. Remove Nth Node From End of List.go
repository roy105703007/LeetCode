/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var behind *ListNode
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	pre := new(ListNode)
	pre.Next = behind
	head = pre
	for behind != nil {
		if n == 1 {
			pre.Next = pre.Next.Next
			break
		} else {
			pre = behind
			behind = behind.Next
			n--
		}
	}
	head = head.Next
	behind = nil
	for head != nil {
		next := head.Next
		head.Next = behind
		behind = head
		head = next
	}
	return behind
}