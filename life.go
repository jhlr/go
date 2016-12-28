package life

type Universe struct {
	r     Rule
	count uint
	board map[Position][2]bool
}

func New(r Rule) *Universe {
	u := new(Universe)
	u.r = r
	u.count = 0
	u.board = make(map[Position][2]bool)
	return u
}

func (u *Universe) Next() {
	c := u.count % 2
	for p := range u.board {
		u.Around(p, func(q Position) {
			b := u.next(q, c)
			u.set(q, 1-c, b)
		})
	}
	u.count++
	u.trim()
}

func (u *Universe) Set(p Position, b bool) {
	u.set(p, u.count%2, b)
}

func (u *Universe) Get(p Position) bool {
	return u.get(p, u.count%2)
}

func (u *Universe) Count() uint {
	return u.count
}

func (u *Universe) Rule() Rule {
	return u.r
}

func (u *Universe) Around(p Position, foo func(Position)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			q := P(p.X+i, p.Y+j)
			foo(q)
		}
	}
}

func (u *Universe) trim() {
	c := u.count % 2
	for p := range u.board {
		if !u.get(p, c) {
			delete(u.board, p)
		}
	}
}

func (u *Universe) get(p Position, c uint) bool {
	return u.board[p][c]
}

func (u *Universe) set(p Position, c uint, b bool) {
	cell := u.board[p]
	cell[c] = b
	u.board[p] = cell
}

func (u *Universe) next(p Position, c uint) bool {
	alive := 0
	u.Around(p, func(q Position) {
		if u.get(q, c) {
			alive++
		}
	})
	if u.get(p, c) {
		alive--
		return u.r[1][alive]
	}
	return u.r[0][alive]
}
