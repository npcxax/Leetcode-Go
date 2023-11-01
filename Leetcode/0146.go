package leetcode

// my solution
// type Node struct {
// 	Key  int
// 	Val  int
// 	Pre  *Node
// 	Next *Node
// }

// type LRUCache struct {
// 	hashmap    map[int]*Node
// 	head, tail *Node
// 	cur, cap   int
// }

// func Constructor(capacity int) LRUCache {
// 	l := LRUCache{hashmap: make(map[int]*Node), head: &Node{}, tail: &Node{}, cur: 0, cap: capacity}
// 	l.head.Next = l.tail
// 	l.tail.Pre = l.head
// 	return l
// }

// func (this *LRUCache) Get(key int) int {
// 	if v, ok := this.hashmap[key]; ok {
// 		if this.tail.Pre == v {
// 			this.deleteTail()
// 		} else {
// 			this.deleteMiddle(v)
// 		}
// 		v := this.addHead(v.Key, v.Val)
// 		this.hashmap[key] = v
// 		return v.Val
// 	}
// 	return -1
// }

// func (this *LRUCache) Put(key int, value int) {
// 	v, ok := this.hashmap[key]
// 	if ok {
// 		if this.tail.Pre == v {
// 			this.deleteTail()
// 		} else {
// 			this.deleteMiddle(v)
// 		}
// 		v = this.addHead(key, value)
// 		this.hashmap[key] = v
// 	} else {
// 		if this.cur == this.cap {
// 			t := this.deleteTail()
// 			delete(this.hashmap, t.Key)
// 			this.cur--
// 		}
// 		this.cur++
// 		v := this.addHead(key, value)
// 		this.hashmap[key] = v
// 	}
// }

// func (this *LRUCache) deleteMiddle(node *Node) *Node {
// 	pre := node.Pre
// 	pre.Next = node.Next
// 	node.Next.Pre = pre
// 	return node
// }

// func (this *LRUCache) deleteTail() *Node {
// 	v := this.tail.Pre
// 	pre := v.Pre
// 	pre.Next = this.tail
// 	this.tail.Pre = pre
// 	return v
// }

// func (this *LRUCache) addHead(key int, value int) *Node {
// 	newNode := &Node{Key: key, Val: value, Pre: this.head, Next: this.head.Next}
// 	this.head.Next.Pre = newNode
// 	this.head.Next = newNode
// 	return newNode
// }

// official
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

func initDLinkNode(key, value int) *DLinkNode {
	return &DLinkNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkNode),
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.moveToHead(v)
		return v.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cache[key]
	if ok {
		v.value = value
		this.moveToHead(v)
	} else {
		if this.size == this.capacity {
			t := this.removeTail()
			delete(this.cache, t.key)
			this.size--
		}
		this.size++
		v = initDLinkNode(key, value)
		this.addToHead(v)
		this.cache[key] = v
	}
}

func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
