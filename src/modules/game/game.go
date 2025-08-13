package game

type Game struct {
	Grid      [][]Cell
	Size      int
	Obstacles []Position
	Commands  string
}

type Cell struct {
	X        int
	Y        int
	HasRover bool
}

type Position struct {
	X int
	Y int
}

func NewGame(size int, obstacles []Position, commands string) *Game {
	grid := make([][]Cell, size)
	for i := range grid {
		grid[i] = make([]Cell, size)
	}
	return &Game{
		Grid:      grid,
		Size:      size,
		Obstacles: obstacles,
		Commands:  commands,
	}
}
