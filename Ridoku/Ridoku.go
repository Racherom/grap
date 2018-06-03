package Ridoku

import (
	"fmt"
)

type Level []int

var (
	Level1 = Level{3, 4, 5, 4, 3}
	Level2 = Level{5, 6, 7, 8, 9, 8, 7, 6, 5}
	Level3 = Level{8, 9, 8, 9, 8, 9, 8, 9, 8, 9, 8}
)

type (
	Board      [][]Tile
	Connection struct {
		Tiles      [2]*Tile
		Successive bool
	}
	Tile struct {
		N     int
		sides [6]*Connection
	}
)

func (l Level) valid() bool {
	for i := range l {
		if i > 0 &&
			!(l[i] == l[i-1]-1 || l[i] == l[i-1]+1) {
			return false
		}
	}
	return true
}

func (t1 *Tile) connectTo(t2 *Tile, side int) {
	if t1.N == -1 || t2.N == -1 {
		return
	}
	c := Connection{Tiles: [2]*Tile{t1, t2}}
	t1.sides[side] = &c
	opside := side + 3
	if side > 2 {
		opside = side - 3
	}
	t2.sides[opside] = &c
}

func NewBoard(level Level) (Board, error) {
	if !level.valid() {
		return Board{}, fmt.Errorf("invalid level")
	}

	b := make(Board, len(level))
	h := len(level) / 2
	hr := level[h] / 2
	for i := 0; i < len(level); i++ {
		b[i] = make([]Tile, level[i])
		if i == h {
			b[h][hr].N = -1
		}
		for j := 0; j < level[i]; j++ {
			t := &b[i][j]

			if j > 0 {
				t.connectTo(&b[i][j-1], 4)
			}

			if i == 0 {
				continue
			}

			if level[i] > level[i-1] {
				if j > 0 {
					t.connectTo(&b[i-1][j-1], 5)
				}
				if j < level[i-1] {
					t.connectTo(&b[i-1][j], 0)
				}
			} else {
				t.connectTo(&b[i-1][j], 5)
				t.connectTo(&b[i-1][j+1], 0)
			}

		}
	}

	return b, nil
}

func (b *Board) get(r, c int) *Tile {
	if len(*b) > r && len((*b)[r]) > c {
		return &(*b)[r][c]
	}
	return nil
}

func (b *Board) getByN(n int) *Tile {
	for i := range *b {
		for j := range (*b)[i] {
			if (*b)[i][j].N == n {
				return &(*b)[i][j]
			}
		}
	}
	return nil
}

func (b *Board) Successive(r, c int, s ...int) error {
	t := b.get(r, c)
	if t == nil {
		return fmt.Errorf("invalid loc")
	}
	suc := true
	for i := range s {
		if s[i] < 6 && t.sides[s[i]] != nil {
			t.sides[s[i]].Successive = true
		} else {
			suc = false
		}
	}
	if !suc {
		return fmt.Errorf("invalide side providet")
	}
	return nil
}

func (b *Board) Set(r, c, n int) error {
	t := b.get(r, c)
	if t != nil {
		t.N = n
		return nil
	}
	return fmt.Errorf("invalid loc")
}

func (b Board) Copy() Board {
	l := make([]int, 0)
	for i := range b {
		l = append(l, len(b[i]))
	}
	bn, _ := NewBoard(l)
	for i := range b {
		for j := range b[i] {
			bn[i][j].N = b[i][j].N
			for k := range b[i][j].sides {
				if b[i][j].sides[k] != nil {
					bn[i][j].sides[k].Successive = b[i][j].sides[k].Successive
				}
			}
		}
	}
	return bn
}

func (bs Board) Slove() Board {
	boards := []Board{bs}
	max := bs.Max()
	for n := 2; n <= max; n++ {
		newBoards := make([]Board, 0)
		isSet := bs.getByN(n) != nil
		fmt.Printf("%d: %d\n sedr ", n, len(boards))
		for i := range boards {
			b := boards[i]
			t := b.getByN(n - 1)
			found := false
			for j := range t.sides {
				if t.sides[j] == nil || found {
					continue
				}
				for k := range t.sides[j].Tiles {
					if t.sides[j].Tiles[k].N == n {
						newBoards = append(newBoards, b)
						found = true
						break
					}
					if t.sides[j].Successive && t.sides[j].Tiles[k].N == 0 {
						t.sides[j].Tiles[k].N = n
						newBoards = append(newBoards, b)
						found = true
						break
					}
				}
			}
			if isSet || found {
				continue
			}
			for j := range t.sides {
				if t.sides[j] == nil {
					continue
				}
				for k := range t.sides[j].Tiles {
					if t.sides[j].Tiles[k].N == 0 {
						t.sides[j].Tiles[k].N = n
						newBoards = append(newBoards, b.Copy())
						t.sides[j].Tiles[k].N = 0
					}
				}
			}
		}
		boards = newBoards

	}
	return boards[0]
}

func (b Board) Print() {
	fmt.Println()
	for i := range b {
		if len(b[i])%2 == 0 {
			fmt.Print("  ")
		}
		for j := range b[i] {
			fmt.Printf("%d;", b[i][j].N)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) Max() int {
	max := -1
	for i := range *b {
		max += len((*b)[i])
	}
	return max
}

func (b *Board) SetCorrecct() {
	for i := range *b {
		for j := range (*b)[i] {
			(*b)[i][j].N = Correct[i][j].N
		}
	}
}

var Correct = Board{
	[]Tile{{N: 76}, {N: 77}, {N: 78}, {N: 79}, {N: 80}, {N: 81}, {N: 82}, {N: 84}},
	[]Tile{{N: 74}, {N: 75}, {N: 67}, {N: 66}, {N: 65}, {N: 57}, {N: 55}, {N: 83}, {N: 85}},
	[]Tile{{N: 73}, {N: 68}, {N: 69}, {N: 64}, {N: 58}, {N: 56}, {N: 54}, {N: 86}},
	[]Tile{{N: 72}, {N: 71}, {N: 70}, {N: 0}, {N: 63}, {N: 59}, {N: 53}, {N: 87}, {N: 88}},
	[]Tile{{N: 8}, {N: 6}, {N: 0}, {N: 0}, {N: 62}, {N: 60}, {N: 52}, {N: 89}},
	[]Tile{{N: 9}, {N: 7}, {N: 12}, {N: 2}, {N: -1}, {N: 61}, {N: 51}, {N: 50}, {N: 90}},
	[]Tile{{N: 10}, {N: 11}, {N: 13}, {N: 1}, {N: 16}, {N: 17}, {N: 49}, {N: 91}},
	[]Tile{{N: 31}, {N: 30}, {N: 29}, {N: 14}, {N: 15}, {N: 19}, {N: 18}, {N: 48}, {N: 92}},
	[]Tile{{N: 32}, {N: 28}, {N: 26}, {N: 25}, {N: 21}, {N: 20}, {N: 46}, {N: 47}},
	[]Tile{{N: 33}, {N: 35}, {N: 27}, {N: 24}, {N: 23}, {N: 22}, {N: 45}, {N: 44}, {N: 43}},
	[]Tile{{N: 34}, {N: 36}, {N: 37}, {N: 38}, {N: 39}, {N: 40}, {N: 41}, {N: 42}},
}
