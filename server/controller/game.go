package controller

import (
	"github.com/gin-gonic/gin"
	"goSumerGame/server/gameplay"
	"goSumerGame/server/model"
	"net/http"
)

func AddGame(context *gin.Context) {
	var input model.Game
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred internally"})
	}

	input.UserID = user.ID

	newGame := gameplay.Game{}
	newGameState := gameplay.GameState{}
	newGame.Initialize(input.Debug, &newGameState)

	err := newGameState.Initialize(input.Debug)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedEntry, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = newGame.Save(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Updating here to reflect the changes to the input.Location field following the creation of the save file itself
	// This seems safer than trying to infer the correct savegame ID by querying for the last ID in the database, and then adding 1 to it.
	// It also is probably more performant if the database ends up really large
	savedEntry, err = input.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func DeleteGame(context *gin.Context) {
	var input model.Game
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user has been set in auth middleware
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred internally"})
	}

	input.UserID = user.ID

	id, err := input.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusFound, gin.H{"data": id})
}

func GetAllGames(context *gin.Context) {
	// user has been set in auth middleware
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured internally"})
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Games})
}
