package rover

import (
	"mars-rover-navigation/src/model"
	"sync"
)

type roverImpl struct {
	Position  model.Position
	Direction model.Direction
}

var (
	instance *roverImpl
	once     sync.Once
)

func NewRover(x, y int, direction model.Direction) *roverImpl {
	once.Do(func() {
		instance = &roverImpl{
			Position:  model.Position{X: x, Y: y},
			Direction: direction,
		}
	})
	return instance
}

func Reset() {
	once = sync.Once{}
	instance = nil
}

func (r *roverImpl) GetTryMovePosition() model.Position {
	expectMove := r.Position

	switch r.Direction {
	case "N":
		expectMove.Y++
	case "S":
		expectMove.Y--
	case "E":
		expectMove.X++
	case "W":
		expectMove.X--
	}

	return expectMove
}

func (r *roverImpl) Move() {
	switch r.Direction {
	case "N":
		r.Position.Y++
	case "S":
		r.Position.Y--
	case "E":
		r.Position.X++
	case "W":
		r.Position.X--
	}
}

func (r *roverImpl) TurnLeft() {
	directions := map[model.Direction]model.Direction{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	r.Direction = directions[r.Direction]
}

func (r *roverImpl) TurnRight() {
	directions := map[model.Direction]model.Direction{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}
	r.Direction = directions[r.Direction]
}

func (r *roverImpl) GetPosition() model.Position {
	return r.Position
}

func (r *roverImpl) GetDirection() model.Direction {
	return r.Direction
}
