package linked_list

type Entry struct {
	key   int
	value int
	next  *Entry
}

type MyHashMap struct {
	buckets []*Entry
	size    int
}

func Constructorii() MyHashMap {
	return MyHashMap{
		buckets: make([]*Entry, 10*10*10*10*10*10),
	}
}

func (r *MyHashMap) HashFunction(key int) int {
	return key % len(r.buckets)
}

func (r *MyHashMap) Put(key int, value int) {
	bucketIndex := r.HashFunction(key)
	entry := r.buckets[bucketIndex]

	if entry == nil {
		r.buckets[bucketIndex] = &Entry{key, value, nil}
		r.size++
		return
	}

	for entry != nil {
		if entry.key == key {
			entry.value = value
			return
		}

		if entry.next == nil {
			entry.next = &Entry{key, value, nil}
			return
		}
		entry = entry.next
	}
}

func (r *MyHashMap) Get(key int) int {
	bucketIndex := r.HashFunction(key)

	entry := r.buckets[bucketIndex]

	for entry != nil {
		if entry.key == key {
			return entry.value
		}
		entry = entry.next
	}
	return -1
}

func (r *MyHashMap) Remove(key int) {
	index := r.HashFunction(key)
	head := r.buckets[index]

	var prev *Entry
	for current := head; current != nil; current = current.next {
		if current.key == key {
			if prev == nil {
				r.buckets[index] = current.next
			} else {
				prev.next = current.next
			}
			r.size--
			return
		}
		prev = current
	}
}
