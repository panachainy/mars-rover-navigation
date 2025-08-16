package environment

import "mars-rover-navigation/src/model"

type Environment interface {
	GetGrid() [][]model.Cell
	CanMove(actorPosition model.Position) bool
}
type CanMoveStatus string

const (
	Success             CanMoveStatus = "Success"
	ObstacleEncountered CanMoveStatus = "Obstacle encountered"
	OutOfBounds         CanMoveStatus = "Out of bounds"
)
