package main

type LRUCache struct {
	head  *node
	tail  *node
	store map[int]*node
	cap   int
}

type node struct {
	prev *node
	next *node
	key  int
	val  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{cap: capacity, store: make(map[int]*node)}
}

func (c *LRUCache) removeFromChain(n *node) {
	if n == c.head {
		c.head = n.next
	}

	if n == c.tail {
		c.tail = n.prev
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
}

func (c *LRUCache) addToChain(n *node) {
	n.prev = nil
	n.next = c.head
	if n.next != nil {
		n.next.prev = n
	}
	c.head = n
	if c.tail == nil {
		c.tail = n
	}
}

func (c *LRUCache) Get(key int) int {
	n, ok := c.store[key]
	if !ok {
		return -1
	}

	c.removeFromChain(n)
	c.addToChain(n)
	return n.val
}

func (c *LRUCache) Put(key int, value int) {
	n, ok := c.store[key]
	if !ok {
		n = &node{val: value, key: key}
		c.store[key] = n
	} else {
		n.val = value
		c.removeFromChain(n)
	}

	c.addToChain(n)
	if len(c.store) > c.cap {
		n = c.tail
		if n != nil {
			c.removeFromChain(n)
			delete(c.store, n.key)
		}
	}
}
