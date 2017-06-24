package src

const (
	JailPenalty = 200
	HotelPrice = 150

)
type Cell interface {
   ChangeWorthBy() int
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

func(emptyCell EmptyCell) ChangeWorthBy() int{
	return 0
}

func(hotelCell HotelCell) ChangeWorthBy() int{
	return  - HotelPrice
}

func(jailCell JailCell) ChangeWorthBy() int{
	return  - JailPenalty
}
