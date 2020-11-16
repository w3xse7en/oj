package problem0146

type Node struct {
	value int
	key   int
	pre   *Node
	next  *Node
}

type LRUCache struct {
	kv   map[int]*Node
	head *Node
	tail *Node
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{kv: make(map[int]*Node, capacity), cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.kv[key]
	if !ok {
		return -1
	}
	this.setHead(v, key, v.value)
	return v.value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.kv[key]
	if !ok {
		newNode := &Node{key: key, value: value}
		newNode.next = this.head
		if this.head != nil {
			this.head.pre = newNode
		}
		this.head = newNode
		if len(this.kv) == 0 {
			this.tail = newNode
		}
		this.kv[key] = newNode
		if len(this.kv) > this.cap {
			pre := this.tail.pre
			if pre != nil {
				pre.next = nil
			}
			this.tail.pre = nil
			delete(this.kv, this.tail.key)
			this.tail = pre

		}
	} else {
		this.setHead(v, key, value)
	}
}
func (this *LRUCache) setHead(v *Node, key, value int) {
	v.value = value
	if v == this.head {
		return
	}
	next := v.next
	pre := v.pre
	if v == this.tail && pre != nil {
		this.tail = pre
	}
	if next != nil {
		next.pre = pre
	}
	if pre != nil {
		pre.next = next
	}
	v.pre = nil
	v.next = this.head
	this.head.pre = v
	this.head = v
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
