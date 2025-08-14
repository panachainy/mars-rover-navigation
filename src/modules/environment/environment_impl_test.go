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
