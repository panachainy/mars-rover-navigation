package model

// TODO: will refactor

type Cell struct {
	Position   Position
	HasRover   bool
	IsObstacle bool
}

type Position struct {
	X int
	Y int
}
