package main

import "fmt"

type Room struct {
	ID       int
	BedCount int
	Type     string
	Status   bool //true for reserved, false for available
	Price    float64
}

var rooms = generateRooms()

func main() {
	var input string
	for input != "exit" {
		fmt.Println("Enter a number")
		fmt.Println("1: Room list")
		fmt.Println("2: Add room")
		fmt.Println("3: Reserve room")
		fmt.Scanln(&input)
		switch input {
		case "1":
			getRoomList()
		case "2":
			addRoom()
		case "3":
			reserveRoom()

		}

	}
}

func getRoomList() {
	for _, value := range rooms {
		fmt.Printf("ID:%d,  Bed Count:%d, Status:%v,  Type:%s,  Price:%.2f\n", value.ID, value.BedCount, value.Status, value.Type, value.Price)
	}
}

func getRoomFormInput() Room {
	room := Room{Status: false}
	fmt.Println("Enter room information line by line (ID, Type, BedCount, Price)")
	fmt.Scanln(&room.ID)
	fmt.Scanln(&room.Type)
	fmt.Scanln(&room.BedCount)
	fmt.Scanln(&room.Price)

	return room
}

func reserveRoom() {
	id := 0
	nights := 0
	personCount := 0
	fmt.Println("Enter the room ID")
	fmt.Scanln(&id)
	var room = getRoom(id)
	if room != nil {
		if room.Status {
			fmt.Println("Room is reserved")
		} else if room.Status == false {
			fmt.Println("Enter reserve information line by line (nights, personCount)")
			fmt.Scanln(&nights)
			fmt.Scanln(&personCount)
			roomPrice, tax, discountAmount, finalPrice := calculateRoomPrice(*room, nights, personCount)
			room.Status = true

			fmt.Printf("Room price: %f, tax: %f, discount: %f, final price: %f \n", roomPrice, tax, discountAmount, finalPrice)
		}
	} else {
		fmt.Println("Room not found")
	}

}

func calculateRoomPrice(room Room, nights int, personCount int) (roomPrice float64, tax float64, discountAmount float64, finalPrice float64) {

	discountPercentage := 0.0
	if nights >= 7 && nights <= 15 {
		discountPercentage = 0.1
	} else if nights >= 15 && nights <= 30 {
		discountPercentage = 0.15
	} else if nights > 30 {
		discountPercentage = 0.2
	}
	switch room.Type {
	case "single":
		roomPrice = float64(nights) * room.Price * float64(personCount) * 1.0
	case "double":
		roomPrice = float64(nights) * room.Price * float64(personCount) * 1.2
	case "standard":
		roomPrice = float64(nights) * room.Price * float64(personCount) * 1.3
	case "suite":
		roomPrice = float64(nights) * room.Price * float64(personCount) * 1.5
	}

	tax = roomPrice * 0.09
	discountAmount = roomPrice * discountPercentage
	finalPrice = roomPrice + tax - discountAmount

	return
}

func getRoom(id int) *Room {
	for i := 0; i < len(rooms); i++ {
		if rooms[i].ID == id {
			return &rooms[i]
		}
	}
	return nil
}

func addRoom() {
	room := getRoomFormInput()
	rooms = append(rooms, room)
	fmt.Println("The room added successfully")
}

func generateRooms() []Room {
	rooms := []Room{}
	rooms = append(rooms, Room{
		ID:       1,
		BedCount: 2,
		Type:     "double",
		Status:   false,
		Price:    650000,
	})
	rooms = append(rooms, Room{
		ID:       2,
		BedCount: 1,
		Type:     "single",
		Status:   false,
		Price:    450000,
	})
	rooms = append(rooms, Room{
		ID:       3,
		BedCount: 2,
		Type:     "suite",
		Status:   false,
		Price:    700000,
	})
	rooms = append(rooms, Room{
		ID:       4,
		BedCount: 1,
		Type:     "single",
		Status:   false,
		Price:    350000,
	})
	rooms = append(rooms, Room{
		ID:       5,
		BedCount: 2,
		Type:     "double",
		Status:   false,
		Price:    900000,
	})
	return rooms
}
