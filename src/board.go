package src

type Board struct {
	cells []Cell
}

func NewBoard(cells []Cell) Board{
	return Board{cells: cells}
}

func (board Board) Move(currentPosition, increment int) Cell {
	newPosition := (currentPosition + increment) % len(board.cells)
	return board.cells[newPosition]
}
