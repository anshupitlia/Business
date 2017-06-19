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
