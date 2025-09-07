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
	RoomType string
	Status string
	Price float64
}

func bookRoom(rooms map[string]HotelRoom, roomNumber string) {
	room, exists := rooms[roomNumber]
	if !exists {
		fmt.Printf("Ошибка: номер %s не существует \n", roomNumber)
		return
	}
		if room.Status == Booked {
		fmt.Printf("Ошибка: номер %s уже забронирован \n", roomNumber)
		return
	}
		if room.Status == Maintance {
		fmt.Printf("Ошибка: номер %s на обслуживании \n", roomNumber)
		return
	}
	room.Status = Booked
	rooms[roomNumber] = room
	fmt.Printf("Номер %s уcпешно забронирован \n", roomNumber)
}

func printRooms(rooms map[string] HotelRoom) {
	for number, room := range rooms {
		fmt.Printf("Номер %s: %s, %s, %.2f рублей/ночь \n", number, room.RoomType, room.Status, room.Price)
	}
}

func main(){
	rooms := map[string]HotelRoom{
		"101":{RoomType: Single, Status: Free, Price:2500},
		"102":{RoomType: Single, Status: Free, Price:2500},
		"201":{RoomType: Double, Status: Free, Price:4900},
		"202":{RoomType: Double, Status: Maintance, Price:5480},
		"301":{RoomType: Single, Status: Free, Price:11100},
		"302":{RoomType: Single, Status: Booked, Price:12000},
	}
	fmt.Println("Начальное состояние номеров:")
	printRooms(rooms)
	fmt.Println()

	bookRoom(rooms, "101")
	bookRoom(rooms, "201")
	bookRoom(rooms, "102")
	bookRoom(rooms, "999")

	fmt.Println()

	fmt.Println("Состояние после бронирования:")

	printRooms(rooms)
}