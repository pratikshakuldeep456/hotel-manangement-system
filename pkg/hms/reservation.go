package hms

import (
	"fmt"
	"time"
)

type Reservation struct {
	ID       int
	Guest    *Guest
	Room     *Room
	CheckIn  time.Time
	CheckOut time.Time
	Status   ReservationStatus
}

func NewReservation(Guest *Guest, Room *Room, CheckIn, CheckOut time.Time) *Reservation {
	return &Reservation{ID: GenerateId(),
		Guest:    Guest,
		Room:     Room,
		CheckIn:  CheckIn,
		CheckOut: CheckOut,
		Status:   Confirmed}
}

func (r *Reservation) CancelReservation() error {
	if r.Status != Confirmed {
		return fmt.Errorf(" booking is not confirmed")
	}
	r.Status = Cancelled
	r.Room.CheckOUT()
	return nil
}
