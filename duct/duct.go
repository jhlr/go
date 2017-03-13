package duct

import (
	"sync"
)

// ErrSendClosedDuct ...
type ErrSendClosedDuct struct{}

func (e ErrSendClosedDuct) Error() string {
	return "send on closed duct"
}

// Rule ...
// should lock if needed
type Rule interface {
	To() int
	From() int
	Len() int
}

// Duct ...
type Duct struct {
	closed bool
	buffer []interface{}
	rule   Rule

	mtx  sync.RWMutex
	cond [2]sync.Cond
}

// NewDuct ...
// preallocates a slice of size N
func NewDuct(rule Rule, n int) *Duct {
	d := &Duct{
		closed: false,
		rule:   rule,
		buffer: make([]interface{}, n),
	}
	for i := range d.cond {
		d.cond[i].L = new(sync.Mutex)
	}
	return d
}

func (d *Duct) Len() int {
	return d.rule.Len()
}

func (d *Duct) Close() {
	d.closed = true
	for i := range d.cond {
		d.cond[i].Broadcast()
	}
}

func (d *Duct) Closed() bool {
	return d.closed
}

// Send ...
// calls TrySend repeatedly
// stores the given element at the slice
// blocks while full
// panics ErrSendClosedDuct
func (d *Duct) Send(e interface{}) {
	d.cond[1].L.Lock()
	defer d.cond[1].L.Unlock()
	for {
		if d.trySend(e) {
			return
		}
		d.cond[1].Wait()
	}
}

// TrySend ...
// calls DuctRule.To for index
// blocks while growing
// returns true if successful
// panics ErrSendClosedDuct
func (d *Duct) trySend(e interface{}) bool {
	i := d.rule.To()
	if d.Closed() {
		panic(ErrSendClosedDuct{})
	} else if i >= 0 {
		// allocate
		if i >= len(d.buffer) {
			temp := make([]interface{}, i+1)
			d.mtx.Lock()
			copy(temp, d.buffer)
			d.buffer = temp
			d.mtx.Unlock()
		}
		// store
		d.buffer[i] = e
		d.cond[0].Signal()
		return true
	}
	return false
}

// Recv ...
// blocks while empty or growing
// calls Duct.TryRecv repeatedly
// returns nil only if closed and empty
func (d *Duct) Recv() (interface{}, bool) {
	d.cond[0].L.Lock()
	defer d.cond[0].L.Unlock()
	for {
		e, ok := d.tryRecv()
		if ok == true {
			return e, true
		} else if d.Closed() {
			return nil, false
		}
		d.cond[0].Wait()
	}
}

// TryRecv ...
// blocks while growing
// calls DuctRule.From for index
// returns nil only if empty
func (d *Duct) tryRecv() (interface{}, bool) {
	i := d.rule.From()
	if i >= 0 {
		d.mtx.RLock()
		e := d.buffer[i]
		d.buffer[i] = nil
		d.mtx.RUnlock()

		d.cond[1].Signal()
		return e, true
	}
	return nil, false
}
