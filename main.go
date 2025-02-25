package main

import (
	"fmt"
	"pratikshakuldeep456/hotel-manangement-system/pkg/hms"
	"pratikshakuldeep456/hotel-manangement-system/pkg/hms/payment"
	"time"
)

func main() {

	hmss := hms.GetInstance()
	Guest1 := hms.NewGuest("pratiksha", "xyz@gmail.com", "00000000")
	Guest2 := hms.NewGuest("kuldeep", "xyz@gmail.com", "00000000")
	hmss.AddGuest(Guest1)
	hmss.AddGuest(Guest2)

	fmt.Println("iofihfg")
	room1 := hms.NewRoom(hms.DELUXE, hms.Available, 100)
	hmss.AddRoom(room1)

	room2 := hms.NewRoom(hms.DELUXE, hms.Available, 100)
	hmss.AddRoom(room2)
	fmt.Println("iofihfg")
	checkInDate := time.Now()
	checkOutDate := checkInDate.AddDate(0, 0, 3)
	//BookRoom
	data, err := hmss.BookRoom(Guest1, room1, checkInDate, checkOutDate)
	fmt.Println(data)
	if err != nil {
		fmt.Print(err)
	}
	hmss.CheckIn(data.ID)

	payment := payment.NewCardPayment()
	hmss.CheckOut(data.ID, payment)

	err = hmss.CancelReservation(data.ID)
	if err != nil {
		fmt.Println(err)
	}

	//res2 := hms.NewReservation(Guest2, room2, checkInDate, checkOutDate)
	data1, err := hmss.BookRoom(Guest2, room2, checkInDate, checkOutDate)
	if err != nil {
		fmt.Print(err)
	}
	err = hmss.CancelReservation(data1.ID)
	if err != nil {
		fmt.Println(err)
	}

}
