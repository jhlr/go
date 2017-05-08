package hash

import (
	"sync"
	"table"
)

// MinSize ...
const MinSize = 47

// node of a nil ended single linked list
type node struct {
	key, elem interface{}
	next      *node
}

// hashTable implements a growing hash map
type hashTable struct {
	mtx   sync.RWMutex
	data  []*node
	hash  table.Hasher
	size  int
	grows bool
}

func (n *node) Set(e interface{}) {
	n.elem = e
}

func (n *node) Elem() interface{} {
	return n.elem
}

func (n *node) Key() interface{} {
	return n.key
}

// New allocates a table.Interface based
// on a linked hash algorithm
func New(size int, h table.Hasher) table.Interface {
	t := &hashTable{size: 0, hash: h}
	if size < 0 {
		size = -size
		t.grows = false
	} else {
		t.grows = true
	}
	if size < MinSize {
		size = MinSize
	}
	t.data = make([]*node, size)
	return t
}

func (t *hashTable) Remove(k interface{}) bool {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	found := false
	find(t, k, func(addr **node) {
		if *addr != nil {
			found = true
			t.size--
			*addr = (*addr).next
		}
	})
	return found
}

func (t *hashTable) Node(k interface{}) table.Node {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	var nd table.Node
	find(t, k, func(addr **node) {
		nd = *addr
	})
	return nd
}

func (t *hashTable) Add(k interface{}) table.Node {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	var nd *node
	find(t, k, func(addr **node) {
		if *addr == nil {
			nd = &node{key: k}
			*addr = nd
			t.size++
		} else {
			nd = *addr
		}
	})
	if t.grows && t.Len() > 2*len(t.data) {
		grow(t)
	}
	return nd
}

func (t *hashTable) Do(f func(table.Node)) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	do(t, f)
}

func (t *hashTable) Len() int {
	return t.size
}

func do(t *hashTable, f func(table.Node)) {
	for i := range t.data {
		nd := t.data[i]
		for nd != nil {
			nxt := nd.next
			f(nd)
			nd = nxt
		}
	}
}

func find(t *hashTable, k interface{}, f func(**node)) {
	h := t.hash.Code(k) %
		uint64(len(t.data))
	var prev *node
	nd := t.data[h]
	for nd != nil {
		if nd.Key() == k {
			break
		}
		prev = nd
		nd = nd.next
	}
	if prev == nil {
		f(&t.data[h])
	} else {
		f(&prev.next)
	}
}

func grow(t *hashTable) {
	l := len(t.data)*3 - 2
	newdata := make([]*node, l)
	do(t, func(temp table.Node) {
		nd := temp.(*node)
		nd.next = nil
		h := t.hash.Code(nd.Key()) %
			uint64(len(newdata))
		if newdata[h] == nil {
			newdata[h] = nd
		} else {
			last := newdata[h]
			for last.next != nil {
				last = last.next
			}
			last.next = nd
		}
	})
	t.data = newdata
}
