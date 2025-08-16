package game

import (
	"mars-rover-navigation/src/model"
	"mars-rover-navigation/src/modules/environment"
	"mars-rover-navigation/src/modules/rover"
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
	FinalPosition  model.Position  `json:"final_position"`
	FinalDirection model.Direction `json:"final_direction"`
	Status         Status          `json:"status"`
}

func NewGame() *gameImpl {
	return &gameImpl{}
}

func (e *gameImpl) NavigateRover(size int, obstacles []model.Position, commands string) Result {
	var env environment.Environment = environment.NewEnvironment(size, obstacles)

	// new rover
	var rover rover.Rover = rover.NewRover(0, 0, "N")

	// try integrate together
	for _, cmd := range commands {
		switch cmd {
		case 'M':
			expectNewPosition := rover.GetTryMovePosition()
			canMoveStatus := env.CanMove(expectNewPosition)

			switch canMoveStatus {
			case environment.Success:
				rover.Move()
			case environment.ObstacleEncountered:
				return Result{
					FinalPosition:  rover.GetPosition(),
					FinalDirection: rover.GetDirection(),
					Status:         StatusObstacleEncountered,
				}
			case environment.OutOfBounds:
				return Result{
					FinalPosition:  rover.GetPosition(),
					FinalDirection: rover.GetDirection(),
					Status:         StatusOutOfBounds,
				}
			}
		case 'L':
			rover.TurnLeft()
		case 'R':
			rover.TurnRight()
		}
	}

	return Result{
		FinalPosition:  rover.GetPosition(),
		FinalDirection: rover.GetDirection(),
		Status:         StatusSuccess,
	}
}
