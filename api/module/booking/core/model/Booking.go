package model

import "gorm.io/gorm"

type Booking struct {
	gorm.Model
	BookingId string
	UserID    uint
}
