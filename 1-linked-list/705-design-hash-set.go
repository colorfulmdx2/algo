package linked_list

import (
	"fmt"
	"hash/fnv"
)

type Node1 struct {
	value int
	next  *Node1
}
type MyHashSet struct {
	buckets []*Node1
	size    int
}

func SetConstructor() MyHashSet {
	return MyHashSet{
		buckets: make([]*Node1, 10*10*10*10*10*10),
	}
}

func (s *MyHashSet) HashFunction(value int) int {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%d", value)))
	return int(h.Sum32())%s.size + 1
}

func (s *MyHashSet) Add(key int) {
	bucketIndex := s.HashFunction(key)
	head := s.buckets[bucketIndex]

	for current := head; current != nil; current = current.next {
		if current.value == key {
			return
		}
	}

	newNode := &Node1{value: key, next: head}
	s.buckets[bucketIndex] = newNode
	s.size++
}

func (s *MyHashSet) Remove(key int) {
	bucketIndex := s.HashFunction(key)
	head := s.buckets[bucketIndex]

	var prev *Node1

	for current := head; current != nil; current = current.next {
		if current.value == key {
			if prev == nil {
				// Remove the head of the list
				s.buckets[bucketIndex] = current.next
			} else {
				// Remove the current node
				prev.next = current.next
			}
			s.size--
			return
		}
		prev = current
	}
}

func (s *MyHashSet) Contains(key int) bool {
	bucketIndex := s.HashFunction(key)
	head := s.buckets[bucketIndex]

	for current := head; current != nil; current = current.next {
		if current.value == key {
			return true
		}
	}
	return false
}
