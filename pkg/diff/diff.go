package diff

import (
	"fmt"
	"log/slog"

	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/logs"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

type FurthestPoint struct {
	X    int
	Path Path
}

func (f FurthestPoint) String() string {
	return fmt.Sprintf("X: %d, Path: %s", f.X, f.Path)
}

type Path []Point

func (p Path) String() string {
	var s string
	for i, point := range p {
		if i == 0 {
			s += point.String()
		} else {
			s += " -> " + point.String()
		}
	}
	return s
}

func Myers(a, b []rune) Changes {
	n, o := len(a), len(b)
	m := n + o

	v := make(map[int]FurthestPoint)

	var path []Point

	logs.Debugln()
depth:
	for d := 0; d < m; d++ {
		for k := -d; k <= d; k += 2 {
			var (
				x     int
				prevK int
			)
			if k == -d || (k != d && v[k-1].X < v[k+1].X) {
				x = v[k+1].X
				prevK = k + 1
				slog.Debug("vertical move", "x", x, "prevK", prevK)
			} else {
				x = v[k-1].X + 1
				prevK = k - 1
				slog.Debug("horizontal move", "x", x, "prevK", prevK)
			}

			y := x - k
			start := Point{X: x, Y: y}
			slog.Debug("start", "start", start)

			currentPath := append([]Point(nil), v[prevK].Path...)
			currentPath = append(currentPath, start)
			slog.Debug("current path", "path", currentPath)

			for x < n && y < o && a[x] == b[y] {
				x++
				y++
				currentPath = append(currentPath, Point{X: x, Y: y})
				slog.Debug("match", "x", x, "y", y, "path", currentPath)
			}

			v[k] = FurthestPoint{X: x, Path: currentPath}
			slog.Debug("furthest point", "k", k, "furthest", v[k])

			printV(v)
			if x >= n && y >= o {
				path = currentPath
				logs.ResultSeparator()
				slog.Debug("found path", "path", path)
				break depth
			}
		}
		logs.Debugln()
	}

	return pathToChanges(a, b, path)
}

func MyersByline(a, b []string) ChangesLines {
	n, o := len(a), len(b)
	m := n + o

	v := make(map[int]FurthestPoint)

	var path []Point

	logs.Debugln()
depth:
	for d := 0; d < m; d++ {
		for k := -d; k <= d; k += 2 {
			var (
				x     int
				prevK int
			)
			if k == -d || (k != d && v[k-1].X < v[k+1].X) {
				x = v[k+1].X
				prevK = k + 1
				slog.Debug("vertical move", "x", x, "prevK", prevK)
			} else {
				x = v[k-1].X + 1
				prevK = k - 1
				slog.Debug("horizontal move", "x", x, "prevK", prevK)
			}

			y := x - k
			start := Point{X: x, Y: y}
			slog.Debug("start", "start", start)

			currentPath := append([]Point(nil), v[prevK].Path...)
			currentPath = append(currentPath, start)
			slog.Debug("current path", "path", currentPath)

			for x < n && y < o && a[x] == b[y] {
				x++
				y++
				currentPath = append(currentPath, Point{X: x, Y: y})
				slog.Debug("match", "x", x, "y", y, "path", currentPath)
			}

			v[k] = FurthestPoint{X: x, Path: currentPath}
			slog.Debug("furthest point", "k", k, "furthest", v[k])

			if x >= n && y >= o {
				path = currentPath
				logs.ResultSeparator()
				break depth
			}
		}
		logs.Debugln()
	}

	return pathLinesToChanges(a, b, path)
}

func printV(v map[int]FurthestPoint) {
	for k, f := range v {
		fmt.Printf("k: %d, furthest: %s\n", k, f)
	}
}
