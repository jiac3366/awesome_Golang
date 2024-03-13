package linklist

func removeElements(head *ListNode, val int) *ListNode {
	d := &ListNode{-1, nil}
	d.Next = head
	cur := d

	// 1. 你看, cur是会变的, 所以你也要考虑 cur是否为 nil
	// for cur.Next != nil {
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
		// 2. 你要把 cur = cur.Next 放在一个 else 的{}里面，因为如果出现连续的要删除的 val, 比如 7,7,7
		// 就不能随意 后移cur，不然就会漏掉了
		// cur = cur.Next
	}

	// 3. 你不能返回 head, 而是返回 d 的 Next，因为遇到7,7,7用例的时候，你的 head 只能指向开头的 7，而不是一个全新的数组
	// return head
	return d.Next
}

// =======================

func removeElements(head *ListNode, val int) *ListNode {
	d := &ListNode{-1, nil}
	new := d
	for head != nil {
		if head.Val != val {
			// concat it to new d
			d.Next = &ListNode{head.Val, nil}
			d = d.Next
		}
		head = head.Next
	}

	return new.Next
}
