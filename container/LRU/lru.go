package LRU

import "container/list"

type LRUCache struct {
	cap int
	ll  *list.List            // 存储 *elemetn(<key, value>)， 方便插入和删除
	m   map[int]*list.Element // 快速根据 key 找到 *element(<key,value>)
}

type element struct {
	key   int
	value int
}

func New(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		ll:  list.New(),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get get the value (will always be positive) of the key if the key exists in the cache,
// otherwise return -1
func (this *LRUCache) Get(key int) int {
	// 如果 key 不存在，返回 -1
	e, ok := this.m[key]
	if !ok {
		return -1
	}

	// 如果 key 已经存在需要将 value 移到最前面, 来保证最后面的元素是最近最少访问的
	this.ll.MoveToFront(e)

	return e.Value.(*element).value
}

// Put Set or insert the value if the key is not already present.
// When the cache reached its capacity,
// it should invalidate the least recently used item before inserting a new item
func (this *LRUCache) Put(key int, value int) {
	// 注意: 若 key 已存在，更新 value，并将 Element 移到 Front
	if e, ok := this.m[key]; ok {
		// 注意: 这里需要更新元素的值
		e.Value.(*element).value = value
		this.ll.MoveToFront(e)
		return
	}

	// 容量足够
	if this.cap > 0 {
		// 注意： cap 是容量，而不是长度，所以是要减 1
		this.cap--
		e := this.ll.PushFront(&element{key: key, value: value})
		this.m[key] = e
		return
	}

	// 容量不够，将最后面的元素移除
	back := this.ll.Back()
	this.ll.Remove(back)
	delete(this.m, back.Value.(*element).key)

	// 将 key,value 放在 List 最前面
	e := this.ll.PushFront(&element{key: key, value: value})
	this.m[key] = e
}
