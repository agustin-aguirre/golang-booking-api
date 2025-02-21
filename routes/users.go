package routes

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprint("Could not parse request data because ", err)})
		return
	}

	err = user.Persist()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprint("Could persist user because ", err)})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprint("Could not parse request data because ", err)})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user. Invalid credentials."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Authetication succeded."})
}

func GetAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not fetch events because: %v\nTry again later.", err)})
		return
	}
	context.JSON(http.StatusOK, users)
}
