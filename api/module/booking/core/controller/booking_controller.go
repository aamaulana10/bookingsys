package controller

import (
	"bookingsys/api/core/database"
	"bookingsys/api/module/booking/core/data"
	"bookingsys/api/module/booking/core/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookingRepo struct {
	Db *gorm.DB
}

func New() *BookingRepo {
	db := database.InitDb()
	db.AutoMigrate(&model.Booking{})
	return &BookingRepo{Db: db}
}

func (repository *BookingRepo) GetBookingList(c *gin.Context) {

	fmt.Println("Calling booking list from data layer")

	var booking []model.Booking
	err := data.GetBookingData(repository.Db, &booking)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, booking)
}
