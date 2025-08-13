package rover

import "mars-rover-navigation/src/model"

type roverImpl struct {
	Position  model.Position
	Direction string
}

func NewRoverImpl(x, y int, direction string) *roverImpl {
	return &roverImpl{
		Position:  model.Position{X: x, Y: y},
		Direction: direction,
	}
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
	directions := map[string]string{
		"N": "W",
		"W": "S",
		"S": "E",
		"E": "N",
	}
	r.Direction = directions[r.Direction]
}

func (r *roverImpl) TurnRight() {
	directions := map[string]string{
		"N": "E",
		"E": "S",
		"S": "W",
		"W": "N",
	}
	r.Direction = directions[r.Direction]
}
