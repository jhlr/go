package duct

type Case struct {
	Duct
	Dir
	Value interface{}
}

func (c *Case) Try() bool {
	switch {
	case c.Dir < 0:
		e, ok := c.Duct.tryRecv()
		if ok {
			c.Value = e
			return true
		}
	case c.Dir > 0:
		ok := c.Duct.trySend(c.Value)
		return ok
	}
	return false
}

type Dir int

const (
	SendDir Dir = 1
	RecvDir Dir = -1
)

func Select(deflt bool, cases []Case) int {
	for {
		for i := range cases {
			if cases[i].Try() {
				return i
			}
		}
		if deflt {
			return -1
		}
	}
}
