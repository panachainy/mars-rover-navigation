package model

type Cell struct {
	Position   Position
	IsObstacle bool
}

type Position struct {
	X int
	Y int
}

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)
