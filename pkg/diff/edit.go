package diff

type EditType int

const (
	Keep EditType = iota
	Insert
	Delete
)

type Edit struct {
	Type  EditType
	Value rune
	PosA  int
	PosB  int
}
