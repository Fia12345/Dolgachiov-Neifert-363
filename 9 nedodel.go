package main 

import "fmt"

const (
	Single = "single"
	Double = "double"
	Suite = "suite"
)

const (
	Free = "free"
	Booked = "booked"
	Maintance = "maintance"
)

type HotelRoom struct {
	RooType string
	Status string
	Price float64
}

func bookRoom(room map[string]HotelRoom, roomNumber string) {
	room, exists := rooms[roomNumber]
	if !exists {
		fmt.Printf("Ошибка: номер %s не существует \n", roomNumber)
		return
	}
		if room.Status == Booked {
		fmt.Printf("Ошибка: номер %s уже забронирован \n", roomNumber)
		return
	}
		if room.Status == maintance {
		fmt.Printf("Ошибка: номер %s на обслуживании \n", roomNumber)
		return
	}
	room.Status = Booked
	rooms[roomNumber] = room
	fmt.printf("Номер %s учпешно забронирован \n", roomNumber)
}

func printRooms(rooms map[strings] HotelRoom) {
	for number, room := range rooms {
		fmt.Printf("Номер %s: %s, %s, %.2f рублей/ночь \n", number, room.RoomsType, room.Status, room.Price)
	}
}

func main(){
	rooms := map[string]HotelRoom{
		"101":{RoomsType: Single, Status: Free, Price:2500},
		//Дальше запиши альцгеймер
	}
	fmt.Println("Начальное состояние номеров:")
	printRooms(rooms)
	fmt.Println

	bookRoom(rooms, "101")
	bookRoom(rooms, "201")
	bookRoom(rooms, "102")
	bookRoom(rooms, "999")

	fmt.Println()

	fmt.Println("Состояние после бронирования:")

	printRooms(rooms)
}