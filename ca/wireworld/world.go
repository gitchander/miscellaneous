package main

// https://ru.wikipedia.org/wiki/Wireworld

// states
const (
	empty        = 0
	electronHead = 1
	electronTail = 2
	conductor    = 3
)

func nextState(state int8) int8 {
	switch state {
	case empty:
		return empty
	case electronHead:
		return electronTail
	case electronTail:
		return conductor
	case conductor:
		{
			// count := gridNearHeads(w, Pt(x, y))
			// if (count == 1) || (count == 2) {
			// 	state = electronHead
			// }
		}
	}
	return 0
}

func byteToState(b byte) (int8, bool) {
	var state int8
	switch b {
	case ' ': // space
		state = empty
	case 'H':
		state = electronHead
	case 't':
		state = electronTail
	case '.':
		state = conductor
	default:
		return 0, false
	}
	return state, true
}

type World struct {
	size         Point
	grids        [][][]int8
	gridIndex    int
	neighborhood []Point
}

func NewWorld(size Point) *World {
	return &World{
		size:         size,
		grids:        makeGrids(2, size),
		neighborhood: MooreNeighborhood(1),
	}
}

func makeGrids(n int, size Point) [][][]int8 {
	grids := make([][][]int8, n)
	for i := range grids {
		grids[i] = makeGrid(size)
	}
	return grids
}

func makeGrid(size Point) [][]int8 {
	grid := make([][]int8, size.X)
	for i := range grid {
		grid[i] = make([]int8, size.Y)
	}
	return grid
}

func (w *World) getGrid() (grid [][]int8) {
	grid = w.grids[w.gridIndex]
	return grid
}

func (w *World) GetState(p Point) (state int8) {
	grid := w.getGrid()
	var (
		x = mod(p.X, w.size.X)
		y = mod(p.Y, w.size.Y)
	)
	state = grid[x][y]
	return state
}

func (w *World) SetState(p Point, state int8) {
	grid := w.getGrid()
	var (
		x = mod(p.X, w.size.X)
		y = mod(p.Y, w.size.Y)
	)
	grid[x][y] = state
}

func (w *World) Next() {

	nextIndex := mod(w.gridIndex+1, len(w.grids))

	grid := w.grids[w.gridIndex]
	gridNext := w.grids[nextIndex]

	for y := 0; y < w.size.Y; y++ {
		for x := 0; x < w.size.X; x++ {
			state := grid[x][y]
			switch state {
			case empty:
				// state = empty
			case electronHead:
				state = electronTail
			case electronTail:
				state = conductor
			case conductor:
				{
					p := Pt(x, y)
					count := w.neighborsWithState(p, electronHead)
					if (count == 1) || (count == 2) {
						state = electronHead
					}
				}
			}
			gridNext[x][y] = state
		}
	}

	w.gridIndex = nextIndex
}

func (w *World) neighborsWithState(p Point, state int8) (count int) {
	for _, d := range w.neighborhood {
		var (
			neighborPos   = p.Add(d)
			neighborState = w.GetState(neighborPos)
		)
		if neighborState == state {
			count++
		}
	}
	return count
}

func mod(x, y int) int {
	m := x % y
	if m < 0 {
		m += y
	}
	return m
}

func (w *World) Range(f func(p Point, state int8) bool) {
	grid := w.getGrid()
	for y := 0; y < w.size.Y; y++ {
		for x := 0; x < w.size.X; x++ {
			var (
				p     = Pt(x, y)
				state = grid[x][y]
			)
			ok := f(p, state)
			if !ok {
				return
			}
		}
	}
}

func ParseLines(w *World, pos Point, lines []string) {
	for y, line := range lines {
		bs := []byte(line)
		for x, b := range bs {
			state, ok := byteToState(b)
			if ok {
				p := Pt(x, y)
				w.SetState(pos.Add(p), state)
			}
		}
	}
}

func ParseFile(w *World, pos Point, filename string) error {
	lines, err := FileLines(filename)
	if err != nil {
		return err
	}
	ParseLines(w, pos, lines)
	return nil
}
