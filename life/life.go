package life

import "image"

// cellsize has to be >=2
const cellsize = 2

type Universe struct {
	rule  Rule
	count uint
	board map[image.Point][cellsize + 1]bool
}

func New(r Rule) *Universe {
	u := new(Universe)
	u.rule = r
	u.count = 0
	u.board = make(map[image.Point][cellsize + 1]bool)
	return u
}

func (u *Universe) Next() {
	s := u.count % cellsize
	for p := range u.board {
		u.Around(p, func(q image.Point) {
			if !u.get(q, cellsize) {
				b := u.foresee(q, s)
				u.set(q, 1-s, b)
			}
		})
	}
	u.count++
	u.Update()
}

func (u *Universe) Around(p image.Point, foo func(image.Point)) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			q := Pt(p.X+i, p.Y+j)
			foo(q)
		}
	}
}

func (u *Universe) Set(p image.Point, b bool) {
	u.set(p, u.count%cellsize, b)
}

func (u *Universe) Get(p image.Point) bool {
	return u.get(p, u.count%cellsize)
}

func (u *Universe) Count() uint {
	return u.count
}

func (u *Universe) Rule() Rule {
	return u.rule
}

func (u *Universe) Update() {
	s := u.count % cellsize
	for p := range u.board {
		if !u.get(p, s) {
			delete(u.board, p)
		} else {
			u.set(p, cellsize, false)
		}
	}
}

func (u *Universe) get(p image.Point, s uint) bool {
	return u.board[p][s]
}

func (u *Universe) set(p image.Point, s uint, b bool) {
	cell := u.board[p]
	cell[cellsize] = true
	cell[s] = b
	u.board[p] = cell
}

func (u *Universe) foresee(p image.Point, s uint) bool {
	alive := 0
	u.Around(p, func(q image.Point) {
		if u.get(q, s) {
			alive++
		}
	})
	if u.get(p, s) {
		alive--
		return u.rule[1][alive]
	}
	return u.rule[0][alive]
}
