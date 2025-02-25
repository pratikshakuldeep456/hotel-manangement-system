package hms

import (
	"fmt"
	"pratikshakuldeep456/hotel-manangement-system/pkg/hms/payment"
	"sync"
	"time"
)

type HotelManagmentSystem struct {
	Rooms        map[int]*Room
	Guests       map[int]*Guest
	Reservations map[int]*Reservation
	// Payment      payment.Payment
}

var Instance *HotelManagmentSystem
var Once sync.Once

func GetInstance() *HotelManagmentSystem {

	Once.Do(func() {
		Instance = &HotelManagmentSystem{
			Guests:       make(map[int]*Guest),
			Rooms:        make(map[int]*Room),
			Reservations: make(map[int]*Reservation),
			//	Payment:      payment,
		}
	})
	return Instance
}

func (hms *HotelManagmentSystem) AddRoom(room *Room) {

	hms.Rooms[room.ID] = room

}
func (hms *HotelManagmentSystem) AddGuest(guest *Guest) {
	hms.Guests[guest.Id] = guest
}

func (hms *HotelManagmentSystem) getRoom(roomId int) error {

	data, exists := hms.Rooms[roomId]
	if !exists {
		return fmt.Errorf("room doesnt exist")
	}
	if data.Status != Available {
		return fmt.Errorf("room is not available for booking")
	}

	return nil

}

func (hms *HotelManagmentSystem) getGuest(guestId int) error {
	_, exists := hms.Guests[guestId]
	if !exists {
		return fmt.Errorf("user doesnt exist")
	}
	return nil
}
func (hms *HotelManagmentSystem) BookRoom(guest *Guest, room *Room, checkInDate, checkOutDate time.Time) (*Reservation, error) {
	err := hms.getRoom(room.ID)
	if err != nil {
		return nil, err
	}

	err = hms.Rooms[room.ID].Book()
	if err != nil {
		return nil, err
	}

	err = hms.getGuest(guest.Id)

	if err != nil {
		return nil, err
	}
	fmt.Println("testing")
	reservation := NewReservation(guest, room, checkInDate, checkOutDate)
	hms.Reservations[reservation.ID] = reservation

	return reservation, nil
}

func (hms *HotelManagmentSystem) CancelReservation(resId int) error {
	data, exists := hms.Reservations[resId]
	if !exists {
		return fmt.Errorf(" reservation doesnt exist")
	}

	if data.Room.Status == Available {
		return fmt.Errorf("cannot cancel reservation after checkout")
	}
	err := data.CancelReservation()
	if err != nil {
		return err
	}

	delete(hms.Reservations, resId)
	return nil
}

func (hms *HotelManagmentSystem) CheckIn(resId int) error {
	data, exists := hms.Reservations[resId]
	if !exists {
		return fmt.Errorf(" reservation doesnt exist")
	}
	if data.Status != Confirmed {
		return fmt.Errorf(" reservation not ")
	}

	return data.Room.CheckIN()

}

func (hms *HotelManagmentSystem) CheckOut(resId int, payment payment.Payment) error {
	data, exists := hms.Reservations[resId]
	if !exists {
		return fmt.Errorf(" reservation doesnt exist")
	}
	if data.Status != Confirmed {
		return fmt.Errorf(" reservation not ")
	}

	days := calculateDays(data.CheckIn, data.CheckOut)
	calculatePrice := days * (float64(data.Room.Price))

	payment.Pay(int(calculatePrice))
	fmt.Println("checkout done ")
	return data.Room.CheckOUT()

}

// func (hms *HotelManagmentSystem) IsAvailable(roomId int, checkin, checkout time.Time) (bool, error) {

// 	data, exists := hms.Rooms[roomId]
// 	if !exists {
// 		return false, fmt.Errorf("room is already booked for selected dates")
// 	}

// 	if data.Status != Available || checkin.Before(time.Now()) || checkin.After(checkout) || checkout.Before(checkin) {
//     return false ,
// 	}
// }
