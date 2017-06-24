package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestEmptyCellDoesNotChangeWorth(t *testing.T) {
	var cell = src.NewEmptyCell()
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != 0 {
		t.Fatalf("Expected 0 but got %d", changeWorthBy)
	}
}

func TestHotelCellDeductsHundredFiftyFromWorth(t *testing.T) {
	var cell = src.NewHotelCell()
	const expectedWorth = src.HotelPrice
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != - expectedWorth {
		t.Fatalf("Expected %d but got %d", - expectedWorth, changeWorthBy)
	}
}

func TestJailCellDeductsTwoHundredFromWorth(t *testing.T) {
	var cell = src.NewJailCell()
	const expectedWorth = src.JailPenalty
	changeWorthBy := cell.ChangeWorthBy()
	if changeWorthBy != - expectedWorth {
		t.Fatalf("Expected %d but got %d", - expectedWorth, changeWorthBy)
	}
}
