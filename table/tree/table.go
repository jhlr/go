package tree

import (
	"sync"
	"table"
)

const MinSize = 0

// treeTable is a synchronized binary avl map.
// It compares its elements using the given
// CompareFunc. It does not make use of the
// reflect package.
type treeTable struct {
	mtx  sync.RWMutex
	comp table.Comparator
	root *node
	size int
}

// New ...
// cmp = func(k,k)int
//   should have deterministic returns
func New(cmp table.Comparator) table.Interface {
	return &treeTable{comp: cmp}
}

// Len returns the size of the map
func (t *treeTable) Len() int {
	return t.size
}

// Add will create a node and Link it to the Map
// using the given key.
func (t *treeTable) Add(k interface{}) table.Node {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	var nd *node

	var add func(*node) *node
	add = func(self *node) *node {
		if self == nil {
			t.size++
			nd = &node{key: k}
			return nd
		}
		i := t.comp.Compare(self.Key(), k)
		if i > 0 {
			self.link[0] = add(self.link[0])
			return self.Fixed()
		} else if i < 0 {
			self.link[1] = add(self.link[1])
			return self.Fixed()
		}
		nd = self
		return self
	}

	t.root = add(t.root)
	return nd
}

func (t *treeTable) Remove(k interface{}) bool {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	found := false
	var remove func(*node) *node
	remove = func(self *node) *node {
		if self == nil {
			return nil
		}
		i := t.comp.Compare(self.Key(), k)
		if i != 0 {
			d := 0
			if i < 0 {
				d = 1
			}
			self.link[d] = remove(self.link[d])
			return self.Fixed()
		}
		found = true
		for i := range self.link {
			if self.link[i] == nil {
				return self.link[1-i]
			}
		}
		var p *node
		n := self.link[1]
		for n.link[0] != nil {
			p = n
			n = n.link[0]
		}
		if p != nil {
			p.link[0] = n.link[1]
		} else {
			self.link[1] = n.link[1]
		}
		n.link = self.link
		return n
	}

	t.root = remove(t.root)
	if found {
		t.size--
	}
	return found
}

// Node will return the node with the given key.
// It will be nil if not found.
func (t *treeTable) Node(k interface{}) table.Node {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	nd := t.root
	for {
		if nd == nil {
			return nil
		}
		i := t.comp.Compare(nd.key, k)
		if i < 0 {
			nd = nd.link[1]
		} else if i > 0 {
			nd = nd.link[0]
		} else {
			return nd
		}
	}
}

// Do will loop through all the nodes of the Map.
// Any operations that modify the map should run
// in another goroutine.
func (t *treeTable) Do(f func(table.Node)) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	var do func(*node)
	do = func(nd *node) {
		for nd != nil {
			if f != nil {
				f(nd)
			}
			do(nd.link[1])
			nd = nd.link[0]
		}
	}
	do(t.root)
}
