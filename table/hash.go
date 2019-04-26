package table

// MinSize ...
const MinSize = 47

// hash implements a growing hash map
type hash struct {
	data  []*nodeh
	hash  Hasher
	size  int
	grows bool
}

// NewHash allocates a Interface based
// on a linked hash algorithm.
// The hashsize will be fixed if the given size
// is negative
func NewHash(size int, h Hasher) Interface {
	t := &hash{size: 0, hash: h}
	if size < 0 {
		size = -size
		t.grows = false
	} else {
		t.grows = true
	}
	if size < MinSize {
		size = MinSize
	} else if size%2 == 0 {
		size++
	}
	t.data = make([]*nodeh, size)
	return t
}

func (t *hash) Remove(k interface{}) bool {
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

func (t *hash) Node(k interface{}) Node {
	var wanted Node
	t.find(k, func(n *nodeh) *nodeh {
		wanted = n
		return n
	})
	return wanted
}

func (t *hash) Add(k interface{}) Node {
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

func (t *hash) Len() int {
	return t.size
}

func (t *hash) Do(f func(Node)) {
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

// find will call f if the node with
// the given key is found
func (t *hash) find(k interface{}, f func(*nodeh) *nodeh) {
	h := t.hash.Hash(k) % uint64(len(t.data))
	var prev *nodeh
	nd := t.data[h]
	for nd != nil {
		if t.hash.Compare(nd.Key(), k) == 0 {
			break
		}
		prev = nd
		nd = nd.next
	}
	if prev == nil {
		t.data[h] = f(t.data[h])
	} else {
		prev.next = f(nd)
	}
}

func (t *hash) grow() {
	l := len(t.data)*3 - 2
	newdata := make([]*nodeh, l)
	t.Do(func(temp Node) {
		nd := temp.(*nodeh)
		// can delete the pointer, it is saved
		nd.next = nil
		h := t.hash.Hash(nd.Key()) % uint64(l)
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
