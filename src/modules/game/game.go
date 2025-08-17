//go:generate go run github.com/golang/mock/mockgen -source=game.go -destination=./mock/mock_game.go -package=mock

package game

import (
	"mars-rover-navigation/src/model"
)

type Game interface {
	NavigateRover(size int, obstacles []model.Position, commands string) Result
}
