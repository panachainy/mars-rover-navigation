package rover

type roverImpl struct {
	X         int
	Y         int
	Direction string
}

func NewRoverImpl(x, y int, direction string) *roverImpl {
	return &roverImpl{
		X:         x,
		Y:         y,
		Direction: direction,
	}
}

func (r *roverImpl) Move() {
	switch r.Direction {
	case "N":
		r.Y++
	case "S":
		r.Y--
	case "E":
		r.X++
	case "W":
		r.X--
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
