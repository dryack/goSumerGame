package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goSumerGame/server/gameplay"
	"goSumerGame/server/model"
	"net/http"
)

// findGameLocation accepts the id of the game being deleted, provided by the
// user, and a pointer to a model.User provided by a call to context.MustGet. It
// searches each Game in the user.Games slice until it finds id. It returns the
// location of the .sav on disk as a string. If no matching id is found, it
// returns an error.
func findGameLocation(id uint, user *model.User) (string, error) {
	for _, game := range user.Games {
		if game.ID == id {
			return game.Location, nil
		}
	}
	return "", errors.New("provided game id not found")
}

func GetAllGames(context *gin.Context) {
	// user has been set in auth middleware
	user, done := userExist(context)
	if !done {
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": user.Games})
}

func userExist(context *gin.Context) (*model.User, bool) {
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred internally"})
		return nil, false
	}
	return user, true
}

func TakeTurn(context *gin.Context) {
	var gameModel model.Game
	var instructions model.Instructions
	if err := context.ShouldBindJSON(&instructions); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gameModel.ID = instructions.GameID

	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred internally"})
		return
	}

	gameSession := gameplay.GameSession{}
	gameModel.UserID = user.ID
	gameId := instructions.GameID
	gameLocation, err := findGameLocation(gameId, user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = gameSession.Load(gameLocation)
	// fmt.Printf("%#v\n%#v\n", gameSession.Meta, gameSession.History[len(gameSession.History)-1]) // debug
	err = gameSession.Test(&instructions, &gameModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%#v\n%#v\n", gameSession.Meta, gameSession.History[len(gameSession.History)-1]) // debug
}
