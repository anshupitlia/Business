package src

import (
	"math/rand"
)

const MaxNumberOnDice = 6

type Player struct {
	position int
	worth int
}

func NewPlayer(position int, worth int) Player {
	return Player {
		position: position,
		worth: worth,
		}
}

func (player Player) Throw() int{
	return rand.Intn(MaxNumberOnDice)
}

func(player Player) Position() int {
	return player.position
}

func(player *Player) MoveTo(position int) {
	player.position = position
}
func(player *Player) ChangeWorth(cell Cell) {
	player.worth = player.worth + cell.ChangeWorthBy()
}

func(player *Player) Change(cell Cell) {
	player.worth = player.worth + cell.ChangeWorthBy()
	player.position = cell.Position()
}

//TODO: this method should not be public
func(player *Player) Worth() int {
	return player.worth
}