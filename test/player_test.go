package test

import (
	"testing"
	"git.periscope-solutions.com/price/Business/src"
)

func TestAPlayerShouldThrowDiceAndGetARandomIncrementBetweenOneAndSix(t *testing.T) {
	var player = src.NewPlayer()
	increment := player.Throw()
	if increment < 1 || increment > src.MaxNumberOnDice {
		t.Fatalf("Expected increment to be between 1 and %d but got %d", src.MaxNumberOnDice, increment)
	}
}