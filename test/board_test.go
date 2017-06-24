package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestBoardShouldFindNextCell(t *testing.T) {
	cell0 := src.NewEmptyCell()
	cell1 := src.NewJailCell()
	cell2 := src.NewHotelCell()
	var cells = []src.Cell{cell0, cell1, cell2}
	var board = src.NewBoard(cells)
	cell := board.Move(0, 2)
	if (cell != cell2) {
		t.Fatalf("Expected %+v but got %+v", cell2, cell)
	}
}
