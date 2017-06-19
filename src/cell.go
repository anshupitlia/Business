package src

const (
	JailPenalty = 200
	HotelPrice = 150

)
type Cell interface {
   changeWorth(int) int
}

type EmptyCell struct {
}

type HotelCell struct {
}

type JailCell struct {
}

func NewEmptyCell() EmptyCell{
	return EmptyCell{}
}

func NewHotelCell() HotelCell {
	return HotelCell{}
}

func NewJailCell() JailCell {
	return JailCell{}
}

func(emptyCell EmptyCell) ChangeWorth(currentWorth int) int{
	return currentWorth
}

func(hotelCell HotelCell) ChangeWorth(currentWorth int) int{
	return currentWorth - HotelPrice
}

func(jailCell JailCell) ChangeWorth(currentWorth int) int{
	return currentWorth - JailPenalty
}
