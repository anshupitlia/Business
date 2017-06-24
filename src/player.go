package src

import "math/rand"

const MaxNumberOnDice = 6

type Player struct {
	position int
}

func NewPlayer(position int) Player {
	return Player{position: position}
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