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
		}
	})

	return instance
}
