package src

import "math/rand"

const MaxNumberOnDice = 6

type Player struct {

}

func NewPlayer() Player {
	return Player{}
}

func (player Player) Throw() int{
	return rand.Intn(MaxNumberOnDice)
}