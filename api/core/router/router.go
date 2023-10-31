package router

import (
	authController "bookingsys/api/module/auth/core/controller"
	bookingController "bookingsys/api/module/booking/core/controller"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	fmt.Print("InitializeRouter")

	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ping")
	})

	bookingRepo := bookingController.New()
	authRepo := authController.New()

	r.GET("/getUsers", authRepo.GetUsersList)
	r.POST("/createUser", authRepo.CreateUser)
	r.GET("/getBookingList", bookingRepo.GetBookingList)

	return r
}
