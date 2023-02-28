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
	var input model.Game
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred internally"})
		return
	}

	gameSession := gameplay.GameSession{}
	input.UserID = user.ID
	gameId := input.Model.ID
	gameLocation, err := findGameLocation(gameId, user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = gameSession.Load(gameLocation)
	fmt.Printf("%#v\n%#v\n", gameSession.Meta, gameSession.History[len(gameSession.History)-1])
	err = gameSession.Test(&input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
