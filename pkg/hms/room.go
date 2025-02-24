package hms

import (
	"fmt"
)

type RoomType string

const (
	SINGLE RoomType = "single"
	DOUBLE RoomType = "double"
	DELUXE RoomType = "deluxe"
	SUITE  RoomType = "suite"
)

type RoomStatus string

const (
	Booked    RoomStatus = "booked"
	Available RoomStatus = "availablee"
	Occupied  RoomStatus = "occupied"
)

type ReservationStatus string

const (
	Confirmed ReservationStatus = "confirmed"
	Cancelled ReservationStatus = "cancelled"
)

type Room struct {
	ID       int
	RoomType RoomType
	Status   RoomStatus

	//ReservationStatus ReservationStatus
	Price int
}

func NewRoom(RoomType RoomType,
	Status RoomStatus,

	//ReservationStatus ReservationStatus,
	Price int) *Room {
	return &Room{ID: GenerateId(),
		RoomType: RoomType,
		Status:   Status,
		//	ReservationStatus: ReservationStatus,
		Price: Price}
}

func (r *Room) IsAvailable() bool {

	return (r.Status == Available)
}

// func (r *Room) IsReservationConfirmed() bool {

// 	return (r.ReservationStatus == Confirmed)

// }

func (r *Room) Book() error {
	if r.Status != Available {
		return fmt.Errorf("room is not available for booking")
	}

	if r.Status == Available {
		r.Status = Booked
	}
	return nil
}

func (r *Room) CheckIN() error {

	if r.Status != Booked {
		return fmt.Errorf("room is not booked so you cant check in")
	}
	//r.CheckIn = time.Now()
	r.Status = Occupied

	return nil

}

func (r *Room) CheckOUT() error {

	if r.Status != Occupied {
		return fmt.Errorf("room is occupied so you cant check out")
	}
	//r.CheckOut = time.Now()
	r.Status = Available

	return nil
}
