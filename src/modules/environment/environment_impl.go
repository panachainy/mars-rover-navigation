package environment

import (
	"mars-rover-navigation/src/model"
)

type environmentImpl struct {
	Size      int
	Obstacles []model.Position
	Grid      [][]model.Cell
}

func NewEnvironment(size int, obstacles []model.Position) *environmentImpl {
	instance := &environmentImpl{
		Size:      size,
		Obstacles: obstacles,
		Grid:      make([][]model.Cell, size),
	}

	for i := range instance.Grid {
		instance.Grid[i] = make([]model.Cell, size)
		for j := range instance.Grid[i] {
			isObstacle := isMatchObstacles(
				model.Position{
					X: i,
					Y: j,
				}, obstacles,
			)
			instance.Grid[i][j] = model.Cell{Position: model.Position{X: i, Y: j}, IsObstacle: isObstacle}
		}
	}
	return instance
}

func (e *environmentImpl) GetGrid() [][]model.Cell {
	return e.Grid
}

func (e *environmentImpl) CanMove(actorPosition model.Position) CanMoveStatus {
	if actorPosition.X < 0 || actorPosition.X >= e.Size || actorPosition.Y < 0 || actorPosition.Y >= e.Size {
		return OutOfBounds
	}

	newRoverGrid := e.Grid[actorPosition.X][actorPosition.Y]

	if newRoverGrid.IsObstacle {
		return ObstacleEncountered
	}

	return Success
}

func isMatchObstacles(position model.Position, obstacles []model.Position) bool {
	for _, o := range obstacles {
		if position.X == o.X && position.Y == o.Y {
			return true
		}
	}

	return false
}
