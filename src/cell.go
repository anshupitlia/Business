package src

type Cell interface {
   ChangeWorthBy() int
}

type EmptyCell struct {
}

type HotelCell struct {
	price int
}

type JailCell struct {
	penalty int
}

func NewEmptyCell() EmptyCell{
	return EmptyCell{}
}

func NewHotelCell(price int) HotelCell {
	return HotelCell{price: price}
}

func NewJailCell(penalty int) JailCell {
	return JailCell{penalty: penalty}
}

func(emptyCell EmptyCell) ChangeWorthBy() int{
	return 0
}

func(hotelCell HotelCell) ChangeWorthBy() int{
	return  - hotelCell.price
}

func(jailCell JailCell) ChangeWorthBy() int{
	return  - jailCell.penalty
}
