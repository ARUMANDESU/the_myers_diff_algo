package diff

import "fmt"

type Table [][]Point

func (t Table) String() string {
	var s string
	for i := range t {
		for j := range t[i] {
			px, py := t[i][j].ParentPos()
			s += fmt.Sprintf("%d(%d:%d) ", t[i][j].Weight, px, py)
		}
		s += "\n"
	}
	return s
}

func (t Table) MaxNeighbour(i, j int) (int, *Point) {
	if i > 0 && j > 0 {
		if t[i-1][j].Weight > t[i][j-1].Weight {
			return t[i-1][j].Weight, &t[i-1][j]
		}
		return t[i][j-1].Weight, &t[i][j-1]
	} else if i > 0 {
		return t[i-1][j].Weight, &t[i-1][j]
	} else if j > 0 {
		return t[i][j-1].Weight, &t[i][j-1]
	}
	return 0, nil
}

type Point struct {
	X, Y   int
	Weight int
	Parent *Point
}

func (p Point) isParent() bool {
	return p.Parent == nil
}

func (p Point) ParentPos() (int, int) {
	if p.Parent == nil {
		return 0, 0
	}

	return p.Parent.X, p.Parent.Y
}
