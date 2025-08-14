package environment

import "mars-rover-navigation/src/model"

type Environment interface {
	GetGrid() [][]model.Cell
}
