package data

import (
	. "bookingsys/api/module/auth/core/model"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, User *User) (err error) {

	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB, User *[]User) (err error) {

	err = db.Find(User).Error
	// err = db.Table("bookings").Error
	if err != nil {
		return err
	}
	return nil
}
