package queue

// Interface is implemented by all queues in
// this package. None of the methods lock.
// The capacity of all queues does not change.
type Interface interface {
	// Add will add an elem. It will return
	// true if queue was already full
	Add(interface{}) bool
	// Rem removes an elem from the queue if
	// possible and returns true if elem is
	// valid
	Rem() (interface{}, bool)
	// Len is the ammount of stored elements
	Len() int
	// Cap is the maximum capacity
	Cap() int
}
