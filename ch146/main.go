// 146. LRU 缓存
// 中等
// 3K
// 相关企业
// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 初始化节点
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// 创建缓存对象
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0), //伪头部
		tail:     initDLinkedNode(0, 0), //伪尾部
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// 获取缓存
func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.cache[key]; !ok {
		return -1
	}
	node := cache.cache[key]
	// 移动节点到链表头部
	cache.moveToHead(node)
	return node.value
}

func (cache *LRUCache) Put(key int, value int) {
	// 判断是否有缓存，如果没有则存入缓存
	if _, ok := cache.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		cache.cache[key] = node
		// 添加到头节点
		cache.addToHead(node)
		cache.size++
		// 如果当前缓存超出容量
		if cache.size > cache.capacity {
			// 移除尾节点
			removed := cache.removeTail()
			delete(cache.cache, removed.key)
			cache.size--
		}
	} else {
		node := cache.cache[key]
		node.value = value
		// 移动到头部
		cache.moveToHead(node)
	}
}

func (cache *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = cache.head
	node.next = cache.head.next
	cache.head.next.prev = node
	cache.head.next = node
}

func (cache *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (cache *LRUCache) moveToHead(node *DLinkedNode) {
	cache.removeNode(node)
	cache.addToHead(node)
}

func (cache *LRUCache) removeTail() *DLinkedNode {
	node := cache.tail.prev
	cache.removeNode(node)
	return node
}

func main() {
	obj := Constructor(2)

	obj.Put(1, 0)

	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	// obj.Put(3, 3)
	// fmt.Println(obj.Get(2))
	// obj.Put(4, 4)
	// fmt.Println(obj.Get(1))
	// fmt.Println(obj.Get(3))
	// fmt.Println(obj.Get(4))
}
