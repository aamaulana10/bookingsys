package data

import (
	"fmt"

	. "bookingsys/api/module/booking/core/model"

	"gorm.io/gorm"
)

func GetBookingData(db *gorm.DB, Booking *[]Booking) (err error) {

	fmt.Println("Calling booking list from database")

	err = db.Find(Booking).Error
	// err = db.Table("bookings").Error
	if err != nil {
		return err
	}
	return nil
}
