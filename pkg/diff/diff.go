package diff

import (
	"log/slog"
)

type Point struct {
	X, Y int
}

type FurthestPoint struct {
	X    int
	Path []Point
}

func Myers(a, b []rune) Changes {
	n, o := len(a), len(b)
	m := n + o

	v := make(map[int]FurthestPoint)

	var path []Point

depth:
	for d := 0; d <= m; d++ {
		for k := -d; k <= d; k += 2 {
			slog.Debug("iteration", "d", d, "k", k)

			var (
				x     int
				prevK int
			)
			if k == -d || (k != d && v[k-1].X < v[k+1].X) {
				x = v[k+1].X // take vertical move
				prevK = k + 1
				slog.Debug("vertical move", "x", x, "prevK", prevK)
			} else {
				x = v[k-1].X + 1 // take horizontal move
				prevK = k - 1
				slog.Debug("horizontal move", "x", x, "prevK", prevK)
			}

			y := x - k
			start := Point{X: x, Y: y}
			slog.Debug("start", "start", start)

			currentPath := append([]Point(nil), v[prevK].Path...)
			currentPath = append(currentPath, start)
			slog.Debug("current path", "path", currentPath)

			for x < len(a) && y < len(b) && a[x] == b[y] {
				x++
				y++
				currentPath = append(currentPath, Point{X: x, Y: y})
				slog.Debug("match", "x", x, "y", y, "path", currentPath)
			}

			v[k] = FurthestPoint{X: x, Path: currentPath}
			slog.Debug("furthest point", "k", k, "furthest", v[k])

			if x >= n && y >= o {
				path = currentPath
				slog.Debug("found path", "path", path)
				break depth
			}
		}
		slog.Debug("")
	}

	return pathToChanges(a, b, path)
}
