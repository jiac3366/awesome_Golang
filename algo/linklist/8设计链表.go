type MyNode struct {
	Val  int
	Next *MyNode
}

type MyLinkedList struct {
	Head *MyNode
	Len  int
}

func Constructor() MyLinkedList {
	head := &MyNode{0, nil}
	return MyLinkedList{
		Head: head,
		Len:  0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Len || index < 0 {
		return -1
	}
	cur := this.Head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmp := this.Head.Next
	newNode := &MyNode{val, tmp}
	this.Head.Next = newNode
	this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Len, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Len || index < 0 {
		return
	}
	cur := this.GotoIndexFront(index)
	tmp := cur.Next
	newNode := &MyNode{val, tmp}
	cur.Next = newNode
	this.Len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Len || index < 0 {
		return
	}
	cur := this.GotoIndexFront(index)
	cur.Next = cur.Next.Next
	this.Len--
}

func (this *MyLinkedList) GotoIndexFront(index int) *MyNode {
	cur := this.Head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
