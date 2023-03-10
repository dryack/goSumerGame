package gameplay

import (
	"testing"
)

func Test_doPlague(t *testing.T) {
	type args struct {
		gameState *GameState
	}
	tests := []struct {
		name                   string
		args                   args
		randomValue            int
		wantString             string
		wantGameStatePlagueVal int
	}{{
		name: "Year 0",
		args: args{
			gameState: &GameState{
				Year: 0,
			},
		},
		randomValue:            5,
		wantString:             "",
		wantGameStatePlagueVal: 0,
	},
		{
			name: "Year 1 no plague",
			args: args{
				gameState: &GameState{
					Year: 1,
				},
			},
			randomValue:            1,
			wantString:             "",
			wantGameStatePlagueVal: 0,
		},
		{
			name: "Year 1 with plague",
			args: args{
				gameState: &GameState{
					Year: 1,
				},
			},
			randomValue:            0,
			wantString:             "A plague has struck!",
			wantGameStatePlagueVal: 1,
		},
		{
			name: "Year 2 without plague",
			args: args{
				gameState: &GameState{
					Year:   2,
					Plague: 1,
				},
			},
			randomValue:            7,
			wantString:             "The plague has ended!",
			wantGameStatePlagueVal: 0,
		},
		{
			name: "Year 2 with plague",
			args: args{
				gameState: &GameState{
					Year:   2,
					Plague: 1,
				},
			},
			randomValue:            0,
			wantString:             "The plague continues!",
			wantGameStatePlagueVal: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := fixedRandomizer{value: tt.randomValue}
			if got := doPlague(tt.args.gameState, r); got != tt.wantString {
				t.Errorf("doPlague() returns = %v, want %v", got, tt.wantString)
			}
			if tt.args.gameState.Plague != tt.wantGameStatePlagueVal {
				t.Errorf("doPlague() gameState.Plague = %v, want %v", tt.args.gameState.Plague, tt.wantGameStatePlagueVal)
			}
		})
	}
}

func Test_doPopulation(t *testing.T) {
	type args struct {
		oldGameState GameState
		newGameState *GameState
		messages     *[]string
	}
	tests := []struct {
		name         string
		args         args
		randomValues []int
		want         int
		wantString   string
	}{
		{
			name: "No bushels released",
			args: args{
				oldGameState: GameState{
					Population: 100,
				},
				newGameState: &GameState{
					BushelsReleased: 0,
				},
				messages: &[]string{},
			},
			randomValues: []int{0, 0, 0, 0},
			want:         50,
			wantString:   "There is starvation in the city!",
		},
		{
			name: "Half required bushels released",
			args: args{
				oldGameState: GameState{
					Population: 100,
				},
				newGameState: &GameState{
					BushelsReleased: 1000,
				},
				messages: &[]string{},
			},
			randomValues: []int{0, 80, 40, 40},
			want:         105,
			wantString:   "There is starvation in the city!",
		},
		{
			name: "All required bushels released",
			args: args{
				oldGameState: GameState{
					Population: 100,
				},
				newGameState: &GameState{
					BushelsReleased: 2000,
				},
				messages: &[]string{},
			},
			randomValues: []int{0, 80, 40, 40},
			want:         155,
			wantString:   "",
		},
		{
			name: "200 pop, no immigrants, emigrants",
			args: args{
				oldGameState: GameState{
					Population: 200,
				},
				newGameState: &GameState{
					BushelsReleased: 4000,
				},
				messages: &[]string{},
			},
			randomValues: []int{2, 0, 0, 0},
			want:         250,
			wantString:   "",
		},
		// TODO: We obviously have a lot of variables we should test, other than the return values
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := sequentialFixedRandomizer{
				called: 0,
				values: tt.randomValues,
			}
			// fmt.Printf("%#v\n", tt.args.newGameState) // debug
			got, got1 := doPopulation(tt.args.oldGameState, tt.args.newGameState, tt.args.messages, &r)
			// fmt.Printf("%#v\n", tt.args.newGameState) // debug
			if got != tt.want {
				t.Errorf("doPopulation() int got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantString {
				t.Errorf("doPopulation() string got = %v, want %v", got1, tt.wantString)
			}
		})
	}
}
