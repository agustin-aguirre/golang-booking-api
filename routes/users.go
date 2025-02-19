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
