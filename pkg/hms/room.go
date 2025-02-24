package hms

import (
	"fmt"
	"time"
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
	ID                int
	RoomType          RoomType
	Status            RoomStatus
	CheckIn           time.Time
	CheckOut          time.Time
	ReservationStatus ReservationStatus
}

func NewRoom() *Room {
	return &Room{}
}

func (r *Room) IsAvailable() bool {
	return (r.Status == Available)
}

func (r *Room) IsReservationConfirmed() bool {

	return (r.ReservationStatus == Confirmed)

}

func (r *Room) Book() {
	if r.Status != Available {
		fmt.Println("room is already occupied or  booked ")
	}

	if r.Status == Available {
		r.Status = Booked
	}
}

func (r *Room) CheckIN() {

	r.CheckIn = time.Now()

}

func (r *Room) CheckOUT() {
	r.CheckIn = time.Now()
}
