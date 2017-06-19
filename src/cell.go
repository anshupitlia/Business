package src

type Cell interface {
   changeWorth(int) int
}

type EmptyCell struct {
}

func NewEmptyCell() EmptyCell{
	return EmptyCell{}
}

func(emptyCell EmptyCell) ChangeWorth(currentWorth int) int{
	return currentWorth
}
