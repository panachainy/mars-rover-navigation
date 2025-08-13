package game

import "mars-rover-navigation/src/model"

type Game struct {
	Commands string
}

func NavigateRover(size int, obstacles []model.Position, commands string) *Game {
	grid := make([][]model.Cell, size)
	for i := range grid {
		grid[i] = make([]model.Cell, size)
	}

	// {
	// "final_position": [0, 4],
	// "final_direction": "N" ,
	// "status": "Out of bounds"
	// }

	return &Game{
		Commands: commands,
	}
}
