package environment

import (
	"mars-rover-navigation/src/model"
	"testing"
)

func TestIsMatchObstacles(t *testing.T) {
	tests := []struct {
		name      string
		position  model.Position
		obstacles []model.Position
		expected  bool
	}{
		{
			name:     "position matches first obstacle",
			position: model.Position{X: 1, Y: 2},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 3, Y: 4},
			},
			expected: true,
		},
		{
			name:     "position matches last obstacle",
			position: model.Position{X: 3, Y: 4},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 3, Y: 4},
			},
			expected: true,
		},
		{
			name:     "position matches middle obstacle",
			position: model.Position{X: 5, Y: 6},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 5, Y: 6},
				{X: 7, Y: 8},
			},
			expected: true,
		},
		{
			name:     "position does not match any obstacle",
			position: model.Position{X: 0, Y: 0},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 3, Y: 4},
			},
			expected: false,
		},
		{
			name:      "empty obstacles list",
			position:  model.Position{X: 1, Y: 1},
			obstacles: []model.Position{},
			expected:  false,
		},
		{
			name:      "nil obstacles list",
			position:  model.Position{X: 1, Y: 1},
			obstacles: nil,
			expected:  false,
		},
		{
			name:     "position with negative coordinates matches obstacle",
			position: model.Position{X: -1, Y: -2},
			obstacles: []model.Position{
				{X: -1, Y: -2},
				{X: 1, Y: 2},
			},
			expected: true,
		},
		{
			name:     "position with zero coordinates matches obstacle",
			position: model.Position{X: 0, Y: 0},
			obstacles: []model.Position{
				{X: 0, Y: 0},
				{X: 1, Y: 1},
			},
			expected: true,
		},
		{
			name:     "position matches X but not Y",
			position: model.Position{X: 1, Y: 3},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 2, Y: 3},
			},
			expected: false,
		},
		{
			name:     "position matches Y but not X",
			position: model.Position{X: 2, Y: 2},
			obstacles: []model.Position{
				{X: 1, Y: 2},
				{X: 3, Y: 3},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchObstacles(tt.position, tt.obstacles)
			if result != tt.expected {
				t.Errorf("isMatchObstacles(%+v, %+v) = %v, expected %v",
					tt.position, tt.obstacles, result, tt.expected)
			}
		})
	}
}

func TestCanMove(t *testing.T) {
	tests := []struct {
		name          string
		size          int
		obstacles     []model.Position
		actorPosition model.Position
		expected      CanMoveStatus
	}{
		{
			name:          "can move to empty cell",
			size:          5,
			obstacles:     []model.Position{{X: 1, Y: 1}, {X: 3, Y: 3}},
			actorPosition: model.Position{X: 0, Y: 0},
			expected:      Success,
		},
		{
			name:          "cannot move to obstacle cell",
			size:          5,
			obstacles:     []model.Position{{X: 1, Y: 1}, {X: 3, Y: 3}},
			actorPosition: model.Position{X: 1, Y: 1},
			expected:      ObstacleEncountered,
		},
		{
			name:          "can move to corner cell",
			size:          3,
			obstacles:     []model.Position{{X: 1, Y: 1}},
			actorPosition: model.Position{X: 2, Y: 2},
			expected:      Success,
		},
		// {
		// 	name:          "cannot move to another obstacle",
		// 	size:          4,
		// 	obstacles:     []model.Position{{X: 0, Y: 1}, {X: 2, Y: 3}},
		// 	actorPosition: model.Position{X: 2, Y: 3},
		// 	expected:      ObstacleEncountered,
		// },
		// {
		// 	name:          "can move when no obstacles",
		// 	size:          3,
		// 	obstacles:     []model.Position{},
		// 	actorPosition: model.Position{X: 1, Y: 1},
		// 	expected:      Success,
		// },
		{
			name:          "can move to edge cell",
			size:          4,
			obstacles:     []model.Position{{X: 1, Y: 1}},
			actorPosition: model.Position{X: 3, Y: 0},
			expected:      Success,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			env := NewEnvironment(tt.size, tt.obstacles)
			result := env.CanMove(tt.actorPosition)
			if result != tt.expected {
				t.Errorf("CanMove(%+v) = %v, expected %v",
					tt.actorPosition, result, tt.expected)
			}
		})
	}
}
