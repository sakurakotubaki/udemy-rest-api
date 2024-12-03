package routes

import (
	"net/http"
	"udemy-rest-api/models"
	"udemy-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
	// 	return
	// }

	// err = user.ValidateCredentials()

	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not validate credentials"})
	// 	return
	// }

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}
