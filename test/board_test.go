package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestBoardShouldFindNextCell(t *testing.T) {
	cell0 := src.NewEmptyCell(0)
	cell1 := src.NewJailCell(1, 200)
	cell2 := src.NewHotelCell(2, 150)
	var cells = []src.Cell{cell0, cell1, cell2}
	var board = src.NewBoard(cells)
	cell := board.Move(0, 2)
	if (cell != cell2) {
		t.Fatalf("Expected %+v but got %+v", cell2, cell)
	}
}

func TestBoardShouldFindNextCellAsHomeCellAfterLastCell(t *testing.T) {
	cell0 := src.NewEmptyCell(0)
	cell1 := src.NewJailCell(1, 200)
	cell2 := src.NewHotelCell(2, 150)
	var cells = []src.Cell{cell0, cell1, cell2}
	var board = src.NewBoard(cells)
	cell := board.Move(2, 1)
	if (cell != cell0) {
		t.Fatalf("Expected %+v but got %+v", cell0, cell)
	}
}
