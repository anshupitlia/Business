package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestAPlayerShouldThrowDiceAndGetARandomIncrementBetweenOneAndSix(t *testing.T) {
	var player = src.NewPlayer(0, 1000)
	increment := player.Throw()
	if increment < 1 || increment > src.MaxNumberOnDice {
		t.Fatalf("Expected increment to be between 1 and %d but got %d", src.MaxNumberOnDice, increment)
	}
}

func TestAPlayerHasACurrentPositionAndReturnsIt(t *testing.T) {
	var player = src.NewPlayer(0, 500)
	position := player.Position()

	if position != 0 {
		t.Fatalf("expected 0 but got %d", position)
	}
}

func TestAPlayerCanBeMovedToAnotherPosition(t *testing.T) {
	var player = src.NewPlayer(0, 1000)
	player.MoveTo(5)
	newPosition := player.Position()

	if newPosition != 5 {
		t.Fatalf("expected 5 but got %d", newPosition)
	}
}

func TestAPlayerChangesItsWorthAndPositionWhenItLandsOnACell(t *testing.T) {
	var player = src.NewPlayer(0, 500)
	var someCell = src.NewHotelCell(1, 150)

	player.Change(someCell)
	newWorth := player.Worth()
	newPosition := player.Position()

	if newWorth != 350 {
		t.Fatalf("expected 350 but got %d", newWorth)
	}

	if newPosition != 1 {
		t.Fatalf("expected to move to position 1 but got %d", newPosition)
	}
}