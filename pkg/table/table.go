package table

const (
	defaultBoardX = 5
	defaultBoardY = 5
)

type Table struct {
	width  int
	height int
}

func New(width int, height int) *Table {
	if width <= 0 {
		width = defaultBoardX
	}
	if height <= 0 {
		height = defaultBoardY
	}
	return &Table{width: width, height: height}
}

func (t *Table) IsValidPosition(x, y int) bool {
	return x >= 0 && x < t.width && y >= 0 && y < t.height
}
