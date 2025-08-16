package console

import (
	"mars-rover-navigation/src/model"
	"mars-rover-navigation/src/modules/game"
	"reflect"
	"testing"
)

func TestConsoleImpl_Success(t *testing.T) {
	grid := 5
	obstacles := []model.Position{{X: 1, Y: 2}, {X: 3, Y: 3}}
	commands := "MMMRM"
	want := game.Result{
		FinalPosition:  model.Position{X: 1, Y: 3},
		FinalDirection: model.East,
		Status:         game.StatusSuccess,
	}

	g := game.NewGame()
	got := g.NavigateRover(grid, obstacles, commands)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestConsoleImpl_ObstacleEncountered(t *testing.T) {
	grid := 5
	obstacles := []model.Position{{X: 1, Y: 2}, {X: 3, Y: 3}}
	commands := "MMRM"
	want := game.Result{
		FinalPosition:  model.Position{X: 0, Y: 2},
		FinalDirection: model.East,
		Status:         game.StatusObstacleEncountered,
	}

	g := game.NewGame()
	got := g.NavigateRover(grid, obstacles, commands)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestConsoleImpl_OutOfBounds(t *testing.T) {
	grid := 5
	obstacles := []model.Position{}
	commands := "MMMMMMMM"
	want := game.Result{
		FinalPosition:  model.Position{X: 0, Y: 4},
		FinalDirection: model.North,
		Status:         game.StatusOutOfBounds,
	}

	g := game.NewGame()
	got := g.NavigateRover(grid, obstacles, commands)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
