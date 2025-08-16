package rover

import "mars-rover-navigation/src/model"

type Rover interface {
	GetTryMovePosition() model.Position

	Move()
	TurnLeft()
	TurnRight()

	GetPosition() model.Position
	GetDirection() model.Direction
}
