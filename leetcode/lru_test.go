package leetcode

// https://leetcode-cn.com/problems/lru-cache/
// https://github.com/hashicorp/golang-lru
func (s *LtCodeSuite) TestP146() {
	// todo not pass

	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	s.Equal(1, l.Get(1))
	s.Equal(-1, l.Get(3))
	l.Put(4, 4) // evicted
	s.Equal(-1, l.Get(2))
	s.Equal(4, l.Get(4))
}

type LRUCache struct {
	capacity int
	items    map[int]int
	indexes  []int
	cur      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		items:    make(map[int]int),
		indexes:  make([]int, capacity),
		cur:      -1,
	}
}

func (l *LRUCache) Get(key int) int {
	if it, ok := l.items[key]; ok {
		l.cur++
		if l.cur == l.capacity {
			l.cur = 0
		}
		l.indexes[l.cur] = key

		return it
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	l.cur++
	if l.cur == l.capacity {
		l.cur = 0
	}

	_, found := l.items[key]
	if found {
		l.items[key] = value
		l.indexes[l.cur] = key
		return
	}

	// not found
	size := len(l.items)
	if size < l.capacity {
		l.items[key] = value
		l.indexes[l.cur] = key
		return
	}

	// evict lru one
	evictKey := l.indexes[l.cur]
	delete(l.items, evictKey)
	l.items[key] = value
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
