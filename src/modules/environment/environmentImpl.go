package environment

import "mars-rover-navigation/src/model"

type environmentImpl struct {
	Grid      [][]model.Cell
	Size      int
	Obstacles []model.Position
}

func (e *environmentImpl) Start() {
}
