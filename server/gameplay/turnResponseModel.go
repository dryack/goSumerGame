package gameplay

type ErrResponse struct {
}

type TurnResponse struct {
	GameState GameState `json:"game_state"`
	Messages  []string  `json:"messages"`
}
