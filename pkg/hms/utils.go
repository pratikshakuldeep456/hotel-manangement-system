package hms

import "time"

var Counter int

func GenerateId() int {
	Counter++
	return Counter
}

func calculateDays(checkIn, checkOut time.Time) float64 {
	days := checkOut.Sub(checkIn).Hours() / 24

	return (days)

}
