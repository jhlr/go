package table

// tree is a synchronized binary avl map.
// It compares its elements using the given
// CompareFunc. It does not make use of the
// reflect package.
type tree struct {
	comp Comparator
	root *nodet
	size int
}

// NewTree ...
// cmp = func(k,k)int
//   should have deterministic returns
func NewTree(cmp Comparator) Interface {
	return &tree{comp: cmp}
}

// Len returns the size of the map
func (t *tree) Len() int {
	return t.size
}

// Add will create or find a node with the given
// key and return it. If list is empty, the given
// key will be compared to itself.
func (t *tree) Add(k interface{}) Node {
	var nd *nodet
	var add func(*nodet) *nodet
	add = func(self *nodet) *nodet {
		if self == nil {
			// found an open leaf
			t.size++
			nd = &nodet{key: k}
			return nd
		}
		i := t.comp.Compare(self.Key(), k)
		if i == 0 {
			nd = self
			return self
		} else if i < 0 {
			self.link[1] = add(self.link[1])
		} else {
			self.link[0] = add(self.link[0])
		}
		return self.fixed()
	}
	if t.root == nil {
		t.comp.Compare(k, k)
	}
	t.root = add(t.root)
	return nd
}

func (t *tree) Remove(k interface{}) bool {
	found := false
	var remove func(*nodet) *nodet
	remove = func(self *nodet) *nodet {
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
			return self.fixed()
		}
		found = true
		for i := range self.link {
			if self.link[i] == nil {
				return self.link[1-i]
			}
		}
		var p *nodet
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
func (t *tree) Node(k interface{}) Node {
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
func (t *tree) Do(f func(Node)) {
	var do func(*nodet)
	do = func(nd *nodet) {
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
