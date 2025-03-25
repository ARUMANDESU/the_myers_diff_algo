package diff

import (
	"fmt"

	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/colors"
	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/logs"
)

// Visualizations have some bugs:
// 1. Not proper alignment of the grid, when the number of characters in the sequence is more than 9
// 2. Move characters are not properly aligned with the grid, especially diagonal and vertical moves

// VisualizeEditGraph creates a visual representation of the edit graph with the path
func VisualizeEditGraph(a, b []rune, path []Point) {
	if !logs.IsDebug() {
		return
	}

	n, m := len(a), len(b)

	// Print header with sequence A
	fmt.Print("         ")
	for i := 0; i < n; i++ {
		fmt.Printf("  %c   ", a[i])
	}
	fmt.Println()

	// Print column indices
	fmt.Print("     ")
	for i := 0; i <= n; i++ {
		fmt.Printf(" %3d  ", i)
	}
	fmt.Println()

	// Create a map for quick lookup if a point is on the path
	pathMap := make(map[string]bool)
	moveType := make(map[string]string) // Store move type for coloring

	// Fill the path map and determine move types
	for i := 1; i < len(path); i++ {
		curr := path[i]
		prev := path[i-1]
		key := fmt.Sprintf("%d,%d", curr.X, curr.Y)
		pathMap[key] = true

		if curr.X == prev.X+1 && curr.Y == prev.Y+1 {
			// Diagonal move (keep)
			moveType[key] = colors.Blue
		} else if curr.X == prev.X+1 {
			// Horizontal move (delete)
			moveType[key] = colors.Red
		} else {
			// Vertical move (insert)
			moveType[key] = colors.Green
		}
	}

	// Add starting point
	if len(path) > 0 {
		startKey := fmt.Sprintf("%d,%d", path[0].X, path[0].Y)
		pathMap[startKey] = true
		moveType[startKey] = colors.Yellow
	}

	// Print each row of the grid
	for j := 0; j <= m; j++ {
		// Print row index
		fmt.Printf("    %d ", j)

		// Print grid cells
		for i := 0; i <= n; i++ {
			key := fmt.Sprintf("%d,%d", i, j)

			if pathMap[key] {
				// This point is on the path, color it based on move type
				color := moveType[key]
				fmt.Printf("%s(%d,%d)%s", color, i, j, colors.Reset)
			} else {
				fmt.Printf("(%d,%d)", i, j)
			}

			// Add connectors between cells based on path
			if i < n {
				// Check if horizontal move is in path
				hKey := fmt.Sprintf("%d,%d", i+1, j)
				if pathMap[key] && pathMap[hKey] {
					fmt.Printf("%s-%s", colors.Red, colors.Reset)
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()

		// Add connectors for vertical and diagonal moves
		if j < m {
			fmt.Printf("%c   ", b[j])
			for i := 0; i <= n; i++ {
				key := fmt.Sprintf("%d,%d", i, j)
				vKey := fmt.Sprintf("%d,%d", i, j+1)
				dKey := fmt.Sprintf("%d,%d", i+1, j+1)

				// Vertical connector
				if pathMap[key] && pathMap[vKey] {
					fmt.Printf(" %s|%s     ", colors.Green, colors.Reset)
				} else {
					fmt.Print("      ")
				}

				// Diagonal connector
				if i < n && pathMap[key] && pathMap[dKey] {
					fmt.Printf(" %s\\%s", colors.Blue, colors.Reset)
				} else if i < n {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}

	// Print legend
	fmt.Println("\nLegend:")
	fmt.Printf("%s--%s Deletion (horizontal move)\n", colors.Red, colors.Reset)
	fmt.Printf("%s|%s  Insertion (vertical move)\n", colors.Green, colors.Reset)
	fmt.Printf("%s\\%s  Match (diagonal move)\n", colors.Blue, colors.Reset)
	fmt.Printf("%s(%d,%d)%s Starting point\n", colors.Yellow, 0, 0, colors.Reset)

	logs.ResultSeparator()
}

// VisualizeProgressiveEditGraph shows the evolution of the edit graph during algorithm execution
func VisualizeProgressiveEditGraph(a, b []rune, v map[int]FurthestPoint, d, k int, start Point) {
	if !logs.IsDebug() {
		return
	}

	fmt.Printf("\nDepth: %d, k-line: %d, Current point: %s\n", d, k, start)

	// Collect all points from all paths in v
	allPoints := make(map[string]struct{})
	for _, fp := range v {
		for _, p := range fp.Path {
			key := fmt.Sprintf("%d,%d", p.X, p.Y)
			allPoints[key] = struct{}{}
		}
	}

	n, m := len(a), len(b)

	space := addSpace(n)

	// Print header
	fmt.Print("         ")
	for i := 0; i < n; i++ {
		if a[i] == ' ' {
			// fmt.Printf("  _   ")
			fmt.Printf("%*s%*c   ", space/2+1, "  ", space/2+1, '_')
		} else {
			fmt.Printf("%*s%*c   ", space/2+1, "  ", space/2+1, a[i])
		}
	}
	fmt.Println()

	fmt.Print("     ")
	for i := 0; i <= n; i++ {
		fmt.Printf("%*s%*d  ", space/2+1, "   ", space/2+1, i)
	}
	fmt.Println()

	// Print grid
	for j := 0; j <= m; j++ {
		fmt.Printf("    %d ", j)

		for i := 0; i <= n; i++ {
			key := fmt.Sprintf("%d,%d", i, j)
			_, inPath := allPoints[key]

			// Highlight current point
			if i == start.X && j == start.Y {
				fmt.Printf("%s(%d,%d)%s", colors.Yellow, i, j, colors.Reset)
			} else if inPath {
				// Show points that have been visited
				fmt.Printf("%s(%d,%d)%s", colors.Blue, i, j, colors.Reset)
			} else {
				fmt.Printf("(%d,%d)", i, j)
			}

			if i < n {
				fmt.Print(" ")
			}
		}
		fmt.Println()

		if j < m {
			fmt.Printf("%c   ", b[j])
			for i := 0; i <= n; i++ {
				fmt.Print("      ")
				if i < n {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}

func addSpace(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	if count%2 == 0 {
		count++
	}

	return count
}
