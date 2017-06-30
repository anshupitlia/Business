package test

import (
	"git.periscope-solutions.com/price/Business/src"
	"testing"
	"reflect"
)

func TestGamePlay(t *testing.T) {
	cell0 := src.NewEmptyCell(0)
	cell1 := src.NewJailCell(1, 200)
	cell2 := src.NewHotelCell(2, 150)
	var cells = []src.Cell{cell0, cell1, cell2}
	var board = src.NewBoard(cells)

	var player1 = src.NewPlayer(0, 1000)
	var player2 = src.NewPlayer(0, 1000)
	game := src.NewGame(board, []src.Player{player1, player2})

	player1 = src.NewPlayer(2, 850)
	player2 = src.NewPlayer(2, 850)
	expectedStateOfGame := src.NewGame(board, []src.Player{player1, player2})
	newStateOfGame := game.Play()

	if (!reflect.DeepEqual(expectedStateOfGame,newStateOfGame)) {
		t.Fatalf("Expected %+v \n but got %+v", expectedStateOfGame, newStateOfGame)
	}
}