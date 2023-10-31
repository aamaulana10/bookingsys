package controller

import (
	"bookingsys/api/core/database"
	"bookingsys/api/module/auth/core/data"
	"bookingsys/api/module/auth/core/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&model.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {

	fmt.Println("Creating user list from data layer")

	var user model.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	err := data.CreateUser(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (repository *UserRepo) GetUsersList(c *gin.Context) {

	fmt.Println("Calling user list from data layer")

	var user []model.User
	err := data.GetUsers(repository.Db, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}
