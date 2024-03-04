package table

const (
	Horizontal         = "\u2500"
	Vertical           = "\u2502"
	TopLeft            = "\u256D"
	TopRight           = "\u256E"
	BottomLeft         = "\u2570"
	BottomRight        = "\u256F"
	VerticalHorizontal = "\u253C"
	VerticalRight      = "\u251C"
	VerticalLeft       = "\u2524"
	DownHorizontal     = "\u252C"
	UpHorizontal       = "\u2534"
)

func New() *TableBuilder {
	return NewTableBuilder()
}
