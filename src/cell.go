package src

type Cell interface {
   ChangeWorthBy() int
	Position() int
}

type EmptyCell struct {
	position int
}

type HotelCell struct {
	position int
	price int
}

type JailCell struct {
	position int
	penalty int
}

func NewEmptyCell(position int) EmptyCell{
	return EmptyCell{position: position}
}

func NewHotelCell(position int ,price int) HotelCell {
	return HotelCell{position: position,
		price: price}
}

func NewJailCell(position int, penalty int) JailCell {
	return JailCell{position: position, penalty: penalty}
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

func(emptyCell EmptyCell) Position() int{
	return emptyCell.position
}

func(hotelCell HotelCell) Position() int{
	return  hotelCell.position
}

func(jailCell JailCell) Position() int{
	return jailCell.position
}


