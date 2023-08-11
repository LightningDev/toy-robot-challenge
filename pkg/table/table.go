package table

var (
	BoardX int = 5
	BoardY int = 5
)

func IsValidPosition(x, y int) bool {
	return x >= 0 && x < BoardX && y >= 0 && y < BoardY
}
