package controller

import (
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

	user, done := userExist(context)
	if !done {
		return
	}

	input.UserID = user.ID

	newGame, err := setupGameSession(input)
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

// setupGameSession initializes the gameplay.GameSession and gameplay.GameState for a new model.Game
func setupGameSession(input model.Game) (gameplay.GameSession, error) {
	newGame := gameplay.GameSession{}
	newGameState := gameplay.GameState{}
	newGame.Initialize(input.Debug, &newGameState)

	err := newGameState.Initialize(input.Debug)
	return newGame, err
}
