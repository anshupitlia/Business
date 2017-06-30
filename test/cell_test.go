package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestEmptyCellDoesNotChangeWorth(t *testing.T) {
	var cell = src.NewEmptyCell(0)
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != 0 {
		t.Fatalf("Expected 0 but got %d", changeWorthBy)
	}
}

func TestHotelCellDeductsHundredFiftyFromWorth(t *testing.T) {
	var cell = src.NewHotelCell(0, 150)
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != -150 {
		t.Fatalf("Expected -150 but got %d", changeWorthBy)
	}
}

func TestJailCellDeductsTwoHundredFromWorth(t *testing.T) {
	var cell = src.NewJailCell(0, 200)
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != -200 {
		t.Fatalf("Expected -200 but got %d", changeWorthBy)
	}
}
