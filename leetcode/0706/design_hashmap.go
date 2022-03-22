package leetcode

type MyHashMap struct {
	m map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		m: make(map[int]int),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	m.m[key] = value
}

func (m *MyHashMap) Get(key int) int {
	v, ok := m.m[key]
	if ok {
		return v
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	delete(m.m, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
