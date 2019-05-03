package table

// MinSize ...
const MinSize = 47

type hashTable struct {
	data  []*nodeh
	hash  Hash
	size  int
	grows bool
}

// NewHash allocates an Interface based on a
// linked hash algorithm. The hashsize will be
// fixed if the given size is negative
func NewHash(size int, h Hash) Interface {
	t := &hashTable{
		size:  0,
		hash:  h,
		grows: true,
	}
	if size < 0 {
		size = -size
		t.grows = false
	}
	if size < MinSize {
		size = MinSize
	} else if size%2 == 0 {
		size++
	}
	t.data = make([]*nodeh, size)
	return t
}

func (t *hashTable) Remove(k interface{}) bool {
	found := false
	t.find(k, func(n *nodeh) *nodeh {
		if n != nil {
			found = true
			t.size--
			return n.next
		}
		return n
	})
	return found
}

func (t *hashTable) Node(k interface{}) Node {
	var wanted Node
	t.find(k, func(n *nodeh) *nodeh {
		if n != nil {
			wanted = n
		}
		return n
	})
	return wanted
}

func (t *hashTable) Add(k interface{}) Node {
	var new *nodeh
	t.find(k, func(n *nodeh) *nodeh {
		if n == nil {
			new = &nodeh{key: k}
			t.size++
		} else {
			new = n
		}
		return new
	})
	if t.grows && t.Len() > 3*len(t.data) {
		t.grow()
	}
	return new
}

func (t *hashTable) Len() int {
	return t.size
}

func (t *hashTable) Do(f func(Node)) {
	for i := range t.data {
		nd := t.data[i]
		for nd != nil {
			nxt := nd.next
			// F can change nd.next
			f(nd)
			nd = nxt
		}
	}
}

// find will call f with the found node
// or nil if not found. K has to have the
// same type and value of the node.
func (t *hashTable) find(k interface{}, f func(*nodeh) *nodeh) {
	h := t.hash.Sum64(k) % uint64(len(t.data))
	var prev *nodeh
	nd := t.data[h]
	for nd != nil {
		if nd.Key() == k {
			// same type and value
			break
		}
		prev = nd
		nd = nd.next
	}
	// either a node or a blank space was found
	if prev == nil {
		t.data[h] = f(t.data[h])
	} else {
		prev.next = f(prev.next)
	}
}

func (t *hashTable) grow() {
	l := len(t.data)*3 - 2
	newdata := make([]*nodeh, l)
	t.Do(func(temp Node) {
		nd := temp.(*nodeh)
		// can delete the pointer, it is saved
		nd.next = nil
		h := t.hash.Sum64(nd.Key()) % uint64(l)
		if newdata[h] == nil {
			// list is empty, just add
			newdata[h] = nd
		} else {
			// add at the end of list
			last := newdata[h]
			for last.next != nil {
				last = last.next
			}
			last.next = nd
		}
	})
	t.data = newdata
}
