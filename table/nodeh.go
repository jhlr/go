package table

// nodeh of a nil ended single linked list
type nodeh struct {
	key, elem interface{}
	next      *nodeh
}

func (n *nodeh) Set(e interface{}) {
	n.elem = e
}

func (n *nodeh) Elem() interface{} {
	return n.elem
}

func (n *nodeh) Key() interface{} {
	return n.key
}
