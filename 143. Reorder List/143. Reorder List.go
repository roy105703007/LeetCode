/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	cur := head
	l := 0
	for cur != nil {
		cur = cur.Next
		l++
	}
	if l < 3 {
		return
	}
	cur = head
	afterMid := 1
	var preMid *ListNode
	for cur.Next != nil {
		if afterMid == l/2 {
			preMid = cur
		}
		if afterMid > l/2 {
			temp := preMid.Next // 原本無
			preMid.Next = cur.Next
			cur.Next = cur.Next.Next
			preMid.Next.Next = temp //原本寫cur，邏輯錯誤，5會連到3

			continue
		}
		cur = cur.Next
		afterMid++
	}
	cur = head
	//這部分以下非常不熟悉
	p := preMid.Next
	for cur != preMid {
		preMid.Next = p.Next
		p.Next = cur.Next
		cur.Next = p
		cur = p.Next
		p = preMid.Next

	}
}