package main

import "fmt"

type Room struct {
	ID       int
	Type     string
	Status   bool
	BedCount int
	Price    int
}

var Rooms []Room = generateRooms()

func main() {

	input := ""
	for input != "exit" {
		fmt.Println("Enter A command:")
		fmt.Println("1: Room Lists")
		fmt.Println("2: Add Room")
		fmt.Println("3: Reserve Room")

		fmt.Scanln(&input)

		switch input {
		case "1":
			getRoomList()
		case "2":
			addRoom()
		case "3":
			reserveRoom()
		case "exit":
			fmt.Println("Exiting....")
			break
		default:
			fmt.Println("Invalid Command!")
		}
	}
}

func getRoomList() {
	for _, room := range Rooms {
		fmt.Printf("%+v\n", room)
	}
}

func getRoomFromInput() Room {
	var room Room = Room{Status: false}
	fmt.Println("Enter room infromation line by line(id, type, badcount, price):")

	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room
}

func addRoom() {
	room := getRoomFromInput()
	Rooms = append(Rooms, room)
}

func reserveRoom() {
	var id int = 0
	fmt.Println("Enter room id for reserver:")
	fmt.Scanln(&id)

	room := GetRoom(id)
	if room == nil {
		fmt.Println("room not found")
		return
	}

	if room.Status {
		fmt.Println("Room already reserve!")
		return
	}

	var nights int = 0
	var personCount int = 0

	fmt.Println("Enter reserve infromation line by line(nights, personCount):")
	fmt.Scanln(&nights)
	fmt.Scanln(&personCount)

	price := calculateRoomPrice(*room, nights, personCount)
	room.Status = true

	fmt.Printf("room price is: %d\n\n", price)

}

func GetRoom(id int) *Room {
	for i := 0; i < len(Rooms); i++ {
		if Rooms[i].ID == id {
			return &Rooms[i]
		}
	}
	return nil
}

func calculateRoomPrice(room Room, nights int, personCount int) (price int) {

	switch room.Type {
	case "Single":
		price = room.Price * nights * personCount

	case "Double":
		price = room.Price * nights * personCount

	case "Family":
		price = room.Price * nights * personCount

	default:
		price = 5000 * nights * personCount
	}

	return price
}

func generateRooms() []Room {
	rooms := []Room{}

	rooms = append(rooms, Room{ID: 1, Type: "Single", Status: true, BedCount: 1, Price: 80})
	rooms = append(rooms, Room{ID: 2, Type: "Double", Status: false, BedCount: 2, Price: 150})
	rooms = append(rooms, Room{ID: 3, Type: "Suite", Status: true, BedCount: 3, Price: 300})
	rooms = append(rooms, Room{ID: 4, Type: "Single", Status: false, BedCount: 1, Price: 70})
	rooms = append(rooms, Room{ID: 5, Type: "Double", Status: true, BedCount: 2, Price: 160})
	rooms = append(rooms, Room{ID: 6, Type: "Family", Status: true, BedCount: 4, Price: 250})
	rooms = append(rooms, Room{ID: 7, Type: "Single", Status: true, BedCount: 1, Price: 90})
	rooms = append(rooms, Room{ID: 8, Type: "Suite", Status: false, BedCount: 3, Price: 280})
	rooms = append(rooms, Room{ID: 9, Type: "Double", Status: true, BedCount: 2, Price: 140})

	return rooms

}
