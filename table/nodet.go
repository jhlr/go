package table

type nodet struct {
	key, elem interface{}
	link      [2]*nodet
	height    int
}

func (n *nodet) Key() interface{} {
	return n.key
}

func (n *nodet) Elem() interface{} {
	return n.elem
}

func (n *nodet) Set(e interface{}) {
	n.elem = e
}

func (n *nodet) fixed() *nodet {
	rotate := func(nd *nodet, d int) *nodet {
		temp := nd.link[1-d]
		nd.link[1-d] = temp.link[d]
		temp.link[d] = nd

		nd.balance()
		temp.balance()
		return temp
	}
	b := n.balance()
	if b < 2 && b > -2 {
		return n
	}
	h := 0
	if b > 0 {
		h = 1
	}
	b *= n.link[h].balance()
	if b < 0 {
		n.link[h] = rotate(n.link[h], h)
	}
	return rotate(n, 1-h)
}

func (n *nodet) balance() int {
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
