package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var slow *ListNode
	fast := head

	for fast != nil {
		tmp := fast
		fast.Next = slow
		slow = fast
		fast = tmp
	}

	return slow
}
