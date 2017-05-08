package tree

type node struct {
	key, elem interface{}
	link      [2]*node
	height    int
}

func (n *node) Key() interface{} {
	return n.key
}

func (n *node) Elem() interface{} {
	return n.elem
}

func (n *node) Set(e interface{}) {
	n.elem = e
}

func (n *node) Fixed() *node {
	rotate := func(nd *node, d int) *node {
		temp := nd.link[1-d]
		nd.link[1-d] = temp.link[d]
		temp.link[d] = nd

		nd.Balance()
		temp.Balance()
		return temp
	}
	b := n.Balance()
	if b < 2 && b > -2 {
		return n
	}
	h := 0
	if b > 0 {
		h = 1
	}
	b *= n.link[h].Balance()
	if b < 0 {
		n.link[h] = rotate(n.link[h], h)
	}
	return rotate(n, 1-h)
}

func (n *node) Balance() int {
	if n == nil {
		return 0
	}
	h := [2]int{0, 0}
	for i, nd := range n.link {
		if nd != nil {
			h[i] = nd.height
		}
	}
	max := h[0]
	if max < h[1] {
		max = h[1]
	}
	n.height = max + 1
	return h[1] - h[0]
}
