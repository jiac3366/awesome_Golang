package linklist

type LinkNode struct {
	Key   int
	Value int
	Prev  *LinkNode
	Next  *LinkNode
}

func NodeConstructor(key, value int) *LinkNode {

	return &LinkNode{
		Key:   key,
		Value: value,
		Prev:  nil,
		Next:  nil,
	}
}

type LRUCache struct {
	Cache    map[int]*LinkNode
	Capacity int
	Head     *LinkNode
	Tail     *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := NodeConstructor(0, 0)
	tail := NodeConstructor(0, 0)
	head.Next = tail
	return LRUCache{
		Cache:    make(map[int]*LinkNode),
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	} else {
		node := this.Cache[key]
		this.MoveToHead(node)
		return node.Value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		// in cache -> update
		node.Value = value
		this.MoveToHead(node)
	} else {
		// add node
		node := NodeConstructor(key, value)
		this.Cache[key] = node
		this.AddHead(node)
		if len(this.Cache) > this.Capacity {
			tail := this.RemoveTail()
			delete(this.Cache, tail.Key)
		}
	}
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddHead(node)
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddHead(node *LinkNode) {
	node.Prev = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
}

func (this *LRUCache) RemoveTail() *LinkNode {
	node := this.Tail.Prev
	this.RemoveNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
