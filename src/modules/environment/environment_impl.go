package environment

import (
	"mars-rover-navigation/src/model"
	"sync"
)

type environmentImpl struct {
	Size      int
	Obstacles []model.Position
	Grid      [][]model.Cell
}

var (
	instance *environmentImpl
	once     sync.Once
)

func NewEnvironment(size int, obstacles []model.Position) *environmentImpl {
	once.Do(func() {
		instance = &environmentImpl{
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
	})

	return instance
}

func Reset() {
	once = sync.Once{}
	instance = nil
}

func (e *environmentImpl) GetGrid() [][]model.Cell {
	return e.Grid
}

func (e *environmentImpl) CanMove(actorPosition model.Position) CanMoveStatus {
	newRoverGrid := e.Grid[actorPosition.X][actorPosition.Y]
	// OutOfBounds

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
