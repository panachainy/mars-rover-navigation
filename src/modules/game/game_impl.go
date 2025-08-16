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
	StatusInvalidInput        Status = "Invalid input"
)

type Result struct {
	FinalPosition  model.Position  `json:"final_position"`
	FinalDirection model.Direction `json:"final_direction"`
	Status         Status          `json:"status"`
}

func NewGame() *gameImpl {
	return &gameImpl{}
}

func isValidInputs(size int, obstacles []model.Position, commands string) bool {
	// Check if size is valid (positive)
	if size <= 0 {
		return false
	}

	// Check if obstacles are within bounds
	for _, obstacle := range obstacles {
		if obstacle.X < 0 || obstacle.X >= size || obstacle.Y < 0 || obstacle.Y >= size {
			return false
		}
	}

	// Check if commands contain only valid characters
	for _, cmd := range commands {
		if cmd != 'L' && cmd != 'R' && cmd != 'M' {
			return false
		}
	}

	return true
}

func (e *gameImpl) NavigateRover(size int, obstacles []model.Position, commands string) Result {

	if !isValidInputs(size, obstacles, commands) {
		return Result{
			FinalPosition:  model.Position{X: 0, Y: 0},
			FinalDirection: model.Direction("N"),
			Status:         StatusInvalidInput,
		}
	}

	var env environment.Environment = environment.NewEnvironment(size, obstacles)
	var rover rover.Rover = rover.NewRover(0, 0, "N")

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
