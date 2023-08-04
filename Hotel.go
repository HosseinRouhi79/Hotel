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

		}
	}
}

func getRoomList() {
	for _, value := range rooms {
		fmt.Printf("ID:%d,  Bed Count:%d,  Type:%s,  Price:%.2f\n", value.ID, value.BedCount, value.Type, value.Price)
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
