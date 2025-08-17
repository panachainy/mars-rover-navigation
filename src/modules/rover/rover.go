//go:generate go run github.com/golang/mock/mockgen -source=rover.go -destination=./mock/mock_rover.go -package=mock

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
