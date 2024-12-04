package routes

import (
	"net/http"
	"strconv"
	"udemy-rest-api/models"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	if event.UserID == userId {
		context.JSON(http.StatusBadRequest, gin.H{"message": "You cannot register for your own event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled successfully"})
}
