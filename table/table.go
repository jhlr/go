package table

// Node is for manipulating the nodes
// and its values. The key is fixed.
// No way of breaking the map using this interface.
type Node interface {
	// Key returns the key of the node.
	// It is not setmaps.
	Key() interface{}
	// Elem returns the held element.
	// It can be set with Set.
	Elem() interface{}
	// Set will replace the held element.
	Set(interface{})
}

// Interface is the respected interface by all the maps
type Interface interface {
	// Add returns the present node or the added one.
	Add(k interface{}) Node
	// Node returns the present node or nil if not found.
	Node(k interface{}) Node
	// Remove will unlink the node with the given key and
	// return whether it was done
	Remove(k interface{}) bool
	// For should call the given function on all the nodes.
	Do(f func(Node))
	Len() int
}

// Comparator is the expected type for comparing keys.
// i == 0 means they are equivalent.
// i < 0 means A should be to the left of B.
// i > 0 means A should be to the right of B.
type Comparator interface {
	Compare(ka, kb interface{}) int
}

// CompareFunc should respect the Comparator interface
type CompareFunc func(ka, kb interface{}) int

// Compare only calls the wrapped CompareFunc
func (f CompareFunc) Compare(ka, kb interface{}) int {
	return f(ka, kb)
}

// Hasher is the interface for hash functions
type Hasher interface {
	Comparator
	Hash(interface{}) uint64
}

// HashFunc returns a 64bit number that ideally only
// represents the given key
type HashFunc func(k interface{}) uint64

// Compare will simply compare the returned values
// from the function calls
func (f HashFunc) Compare(ka, kb interface{}) int {
	ha := f(ka)
	hb := f(kb)
	switch {
	case ha > hb:
		return 1
	case ha < hb:
		return -1
	}
	return 0
}

// Hash will return the result of the call
func (f HashFunc) Hash(k interface{}) uint64 {
	return f(k)
}

// KeyError should be panicked by the CompareFunc
// if any of the given keys is inadequate
type KeyError struct{}

func (e KeyError) Error() string {
	return "given key is inadequate"
}
