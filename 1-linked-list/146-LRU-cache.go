package linked_list

type CacheItem struct {
	key, val   int
	prev, next *CacheItem
}

type LRUCache struct {
	store      map[int]*CacheItem
	capacity   int
	lastResent *CacheItem
	mostResent *CacheItem
}

func constructor(capacity int) LRUCache {
	c := LRUCache{
		store:      make(map[int]*CacheItem),
		capacity:   capacity,
		lastResent: &CacheItem{},
		mostResent: &CacheItem{},
	}

	c.lastResent.next = c.mostResent
	c.mostResent.prev = c.lastResent

	return c
}

func (this *LRUCache) remove(node *CacheItem) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *CacheItem) {
	prevMostRecent, mostRecent := this.mostResent.prev, this.mostResent
	prevMostRecent.next = node
	mostRecent.prev = node
	node.next = mostRecent
	node.prev = prevMostRecent
	// left <-> ... <-> prev <-> right
	// left <-> ... <-> prev <-> node <-> right

}

func (this *LRUCache) Get(key int) int {
	if cacheItem, ok := this.store[key]; ok {
		this.remove(cacheItem)
		this.insert(cacheItem)
		return cacheItem.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// if exists delete node
	// insert node
	// delete last recent is store.len > capacity
	if val, ok := this.store[key]; ok {
		this.remove(val)
		delete(this.store, key)
	}

	node := &CacheItem{
		key: key,
		val: value,
	}
	this.insert(node)
	this.store[key] = node

	if len(this.store) > this.capacity {
		lastRecent := this.lastResent.next
		delete(this.store, lastRecent.key)
		this.remove(lastRecent)
	}
}
