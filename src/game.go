package src

type Game struct {
	board Board
	players []Player
}

func NewGame(board Board, players []Player) Game {
	return Game{
		board: board,
		players: players,
	}
}

func (game Game) Play() Game{
	for i:=0 ; i < len(game.players) ; i++ {
		player := &game.players[i]
		incr := 2 // use player.Throw()
		newCell := game.board.Move(player.position, incr)
		player.Change(newCell)
	}
	return game
}