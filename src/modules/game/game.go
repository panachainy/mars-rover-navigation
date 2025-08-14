package game

import (
	"mars-rover-navigation/src/model"
)

type Game interface {
	NavigateRover(size int, obstacles []model.Position, commands string) Result
}
