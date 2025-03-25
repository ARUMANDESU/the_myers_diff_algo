package diff

import (
	"fmt"
	"strings"

	"github.com/ARUMANDESU/the_myers_diff_algo/pkg/colors"
)

type EditType int

const (
	Keep EditType = iota
	Insert
	Delete
)

type Changes []Edit

func (c Changes) String() string {
	var result strings.Builder

	for _, edit := range c {
		result.WriteString(edit.String())
		result.WriteString("\n")
	}

	return result.String()
}

func (c Changes) ColorString() string {
	var result strings.Builder

	for _, edit := range c {
		result.WriteString(edit.ColorString())
		result.WriteString("\n")
	}

	return result.String()
}

type Edit struct {
	Type  EditType
	Value rune
	PosA  int
	PosB  int
}

func (e Edit) String() string {
	switch e.Type {
	case Keep:
		return fmt.Sprintf(" %c", e.Value)
	case Insert:
		return fmt.Sprintf("+%c", e.Value)
	case Delete:
		return fmt.Sprintf("-%c", e.Value)
	default:
		return "?"
	}
}

func (e Edit) ColorString() string {
	switch e.Type {
	case Keep:
		return fmt.Sprintf("   %c", e.Value)
	case Insert:
		return fmt.Sprintf("%s+++%c%s", colors.Green, e.Value, colors.Reset)
	case Delete:
		return fmt.Sprintf("%s---%c%s", colors.Red, e.Value, colors.Reset)
	default:
		return "?"
	}
}

func FormatEdits(edits []Edit) string {
	var result strings.Builder

	for _, edit := range edits {
		result.WriteString(edit.String())
	}

	return result.String()
}

func pathToChanges(a, b []rune, path []Point) Changes {
	var edits []Edit

	for i := 1; i < len(path); i++ {
		curr := path[i]
		prev := path[i-1]

		if curr.X == prev.X+1 && curr.Y == prev.Y+1 {
			// Diagonal move (keep)
			edits = append(edits, Edit{
				Type:  Keep,
				Value: a[prev.X],
				PosA:  prev.X,
				PosB:  prev.Y,
			})
		} else if curr.X == prev.X+1 {
			// Horizontal move (delete)
			edits = append(edits, Edit{
				Type:  Delete,
				Value: a[prev.X],
				PosA:  prev.X,
				PosB:  prev.Y,
			})
		} else {
			// Vertical move (insert)
			edits = append(edits, Edit{
				Type:  Insert,
				Value: b[prev.Y],
				PosA:  prev.X,
				PosB:  prev.Y,
			})
		}
	}

	return edits
}
