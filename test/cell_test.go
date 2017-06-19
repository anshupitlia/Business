package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestEmptyCellDoesNotChangeWorth(t *testing.T) {
	var cell = src.NewEmptyCell()
	const currentWorth = 500
	changedWorth := cell.ChangeWorth(currentWorth)
	if changedWorth != currentWorth {
		t.Fatalf("Expected %d but got %d", currentWorth, changedWorth)
	}
}
func TestHotelCellDeductsHundredFiftyFromWorth(t *testing.T) {
	var cell = src.NewHotelCell()
	const currentWorth = 500
	const expectedWorth = 350
	changedWorth := cell.ChangeWorth(currentWorth)
	if changedWorth != expectedWorth {
		t.Fatalf("Expected %d but got %d", expectedWorth, changedWorth)
	}
}

func TestJailCellDeductsTwoHundredFromWorth(t *testing.T) {
	var cell = src.NewJailCell()
	const currentWorth = 500
	const expectedWorth = 300
	changedWorth := cell.ChangeWorth(currentWorth)
	if changedWorth != expectedWorth {
		t.Fatalf("Expected %d but got %d", expectedWorth, changedWorth)
	}
}
