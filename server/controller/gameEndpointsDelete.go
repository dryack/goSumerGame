package controller

import (
	"github.com/gin-gonic/gin"
	"goSumerGame/server/gameplay"
	"goSumerGame/server/model"
	"net/http"
)

func DeleteGame(context *gin.Context) {
	var input model.Game
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user has been set in auth middleware
	user, done := userExist(context)
	if !done {
		return
	}

	// determine the location of the physical .sav file
	input.UserID = user.ID
	gameId := input.Model.ID
	gameLocation, err := findGameLocation(gameId, user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameSession := gameplay.GameSession{}
	// delete the .sav file
	err = gameSession.Delete(gameLocation)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := input.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusFound, gin.H{"data": id})
}
