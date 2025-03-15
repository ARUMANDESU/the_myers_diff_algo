package diff

import (
	"fmt"
)

func Myers(a, b []rune) []Edit {
	t := make(Table, len(a))
	for i := range t {
		t[i] = make([]Point, len(b))
		for j := range t[i] {
			t[i][j].X = i
			t[i][j].Y = j
		}
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			t[i][j].Weight, t[i][j].Parent = t.MaxNeighbour(i, j)
			if b[j] == a[i] {
				t[i][j].Weight++
			}
		}
	}

	// slog.Debug("table", "table", points)
	fmt.Println(t)

	return nil
}
