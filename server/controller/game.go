package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"goSumerGame/server/gameplay"
	"goSumerGame/server/model"
	"net/http"
)

// TODO: if the saving of the .sav to disk fails, this should NOT retain the
// created game in the database, and should make it clear that the game creation
// has failed
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

	newGame := gameplay.GameSession{}
	newGameState := gameplay.GameState{}
	newGame.Initialize(input.Debug, &newGameState)

	err := newGameState.Initialize(input.Debug)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = input.Save()
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
	var savedEntry *model.Game
	savedEntry, err = input.Update()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Testing of gameplay.GameSession.Load() functionality
	// testGame := gameplay.GameSession{}
	// testGame.Load(savedEntry)
	// fmt.Printf("%#v", testGame)
	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

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

func GetAllGames(context *gin.Context) {
	// user has been set in auth middleware
	user, ok := context.MustGet("user").(*model.User)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured internally"})
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Games})
}
