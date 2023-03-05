package gameplay

// Instructions represent the player's instructions for a game turn
type Instructions struct {
	GameID            uint
	PurchaseAcres     int
	ReleaseBushels    int
	PlantAcres        int
	PurchaseGranaries int
}
