package diff

import (
	"fmt"
	"strings"
)

type ChangesLines []EditLine

func (c ChangesLines) String() string {
	var result strings.Builder

	for _, edit := range c {
		result.WriteString(edit.String())
		result.WriteString("\n")
	}

	return result.String()
}

type EditLine struct {
	Type EditType
	Line string
	PosA int
	PosB int
}

func (e EditLine) String() string {
	switch e.Type {
	case Keep:
		return fmt.Sprintf(" %s", e.Line)
	case Insert:
		return fmt.Sprintf("+%s", e.Line)
	case Delete:
		return fmt.Sprintf("-%s", e.Line)
	default:
		return "?"
	}
}

func pathLinesToChanges(a, b []string, path []Point) []EditLine {
	edits := make([]EditLine, 0)

	for i := 1; i < len(path); i++ {
		prev := path[i-1]
		curr := path[i]

		if prev.X == curr.X-1 && prev.Y == curr.X-1 {
			edits = append(edits, EditLine{
				Type: Keep,
				Line: a[prev.X],
				PosA: prev.X,
				PosB: prev.Y,
			})
		} else if prev.X == curr.X-1 {
			edits = append(edits, EditLine{
				Type: Delete,
				Line: a[prev.X],
				PosA: prev.X,
				PosB: prev.Y,
			})
		} else {
			edits = append(edits, EditLine{
				Type: Insert,
				Line: b[prev.Y],
				PosA: prev.X,
				PosB: prev.Y,
			})
		}
	}

	return edits
}
