package game

import (
	"mars-rover-navigation/src/model"
	"mars-rover-navigation/src/modules/environment"
	"sync"
)

type gameImpl struct {
}

type Status string

const (
	StatusSuccess             Status = "Success"
	StatusObstacleEncountered Status = "Obstacle encountered"
	StatusOutOfBounds         Status = "Out of bounds"
)

type Result struct {
	FinalPosition  model.Position `json:"final_position"`
	FinalDirection string         `json:"final_direction"`
	Status         Status         `json:"status"`
}

var (
	instance *gameImpl
	once     sync.Once
)

func NewGame() *gameImpl {
	once.Do(func() {
		instance = &gameImpl{}
	})
	return instance
}

func (e *gameImpl) NavigateRover(size int, obstacles []model.Position, commands string) Result {

	env := environment.NewEnvironment(size, obstacles)
	grid := env.GetGrid()

	// FIXME: change mock to real one
	return Result{
		FinalPosition:  model.Position{X: 0, Y: 4},
		FinalDirection: "N",
		Status:         StatusSuccess,
	}
}
