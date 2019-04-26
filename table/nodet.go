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
	if n == nil {
		return nil
	}
	b := n.balance()
	if b < 2 && b > -2 {
		return n
	}

	h := 0
	if b >= 2 {
		h = 1
	}
	b *= n.link[h].balance()
	if b < 0 {
		n.link[h] = n.link[h].rotated(h)
	}
	return n.rotated(1 - h)
}

func (n *nodet) rotated(d int) *nodet {
	if n == nil {
		return nil
	}
	temp := n.link[1-d]
	n.link[1-d] = temp.link[d]
	temp.link[d] = n

	n.balance()
	temp.balance()
	return temp
}

func (n *nodet) balance() int {
	if n == nil {
		return 0
	}
	h := [2]int{0, 0}
	for i := range n.link {
		if n.link[i] != nil {
			h[i] = n.link[i].height
		}
	}
	max := h[0]
	if max < h[1] {
		max = h[1]
	}
	n.height = max + 1
	return h[1] - h[0]
}
