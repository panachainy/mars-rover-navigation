package rover

import (
	"mars-rover-navigation/src/model"
	"testing"
)

func TestNewRover(t *testing.T) {
	tests := []struct {
		name      string
		x         int
		y         int
		direction model.Direction
	}{
		{"Create rover at origin facing North", 0, 0, model.North},
		{"Create rover at positive coordinates facing East", 5, 3, model.East},
		{"Create rover facing West", 10, 20, model.West},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(tt.x, tt.y, tt.direction)

			if rover == nil {
				t.Fatal("NewRover returned nil")
			}

			if rover.Position.X != tt.x {
				t.Errorf("Expected X position %d, got %d", tt.x, rover.Position.X)
			}

			if rover.Position.Y != tt.y {
				t.Errorf("Expected Y position %d, got %d", tt.y, rover.Position.Y)
			}

			if rover.Direction != tt.direction {
				t.Errorf("Expected direction %s, got %s", tt.direction, rover.Direction)
			}
		})
	}
}

func TestGetTryMovePosition(t *testing.T) {
	tests := []struct {
		name      string
		x         int
		y         int
		direction model.Direction
		expectedX int
		expectedY int
	}{
		{"North movement", 5, 5, model.North, 5, 6},
		{"South movement", 5, 5, model.South, 5, 4},
		{"East movement", 5, 5, model.East, 6, 5},
		{"West movement", 5, 5, model.West, 4, 5},
		{"North from origin", 0, 0, model.North, 0, 1},
		{"South from origin", 0, 0, model.South, 0, -1},
		{"East from origin", 0, 0, model.East, 1, 0},
		{"West from origin", 0, 0, model.West, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(tt.x, tt.y, tt.direction)

			tryMovePos := rover.GetTryMovePosition()

			if tryMovePos.X != tt.expectedX {
				t.Errorf("Expected X position %d, got %d", tt.expectedX, tryMovePos.X)
			}

			if tryMovePos.Y != tt.expectedY {
				t.Errorf("Expected Y position %d, got %d", tt.expectedY, tryMovePos.Y)
			}

			// Original position should remain unchanged
			if rover.Position.X != tt.x || rover.Position.Y != tt.y {
				t.Error("GetTryMovePosition should not modify rover's actual position")
			}
		})
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		name      string
		x         int
		y         int
		direction model.Direction
		expectedX int
		expectedY int
	}{
		{"Move North", 5, 5, model.North, 5, 6},
		{"Move South", 5, 5, model.South, 5, 4},
		{"Move East", 5, 5, model.East, 6, 5},
		{"Move West", 5, 5, model.West, 4, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(tt.x, tt.y, tt.direction)

			rover.Move()

			if rover.Position.X != tt.expectedX {
				t.Errorf("Expected X position %d, got %d", tt.expectedX, rover.Position.X)
			}

			if rover.Position.Y != tt.expectedY {
				t.Errorf("Expected Y position %d, got %d", tt.expectedY, rover.Position.Y)
			}

			// Direction should remain unchanged
			if rover.Direction != tt.direction {
				t.Errorf("Direction should remain %s after move", tt.direction)
			}
		})
	}
}

func TestTurnLeft(t *testing.T) {
	tests := []struct {
		name              string
		initialDirection  model.Direction
		expectedDirection model.Direction
	}{
		{"Turn left from North", model.North, model.West},
		{"Turn left from West", model.West, model.South},
		{"Turn left from South", model.South, model.East},
		{"Turn left from East", model.East, model.North},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(0, 0, tt.initialDirection)
			initialPos := rover.Position

			rover.TurnLeft()

			if rover.Direction != tt.expectedDirection {
				t.Errorf("Expected direction %s, got %s", tt.expectedDirection, rover.Direction)
			}

			// Position should remain unchanged
			if rover.Position != initialPos {
				t.Error("TurnLeft should not change rover position")
			}
		})
	}
}

func TestTurnRight(t *testing.T) {
	tests := []struct {
		name              string
		initialDirection  model.Direction
		expectedDirection model.Direction
	}{
		{"Turn right from North", model.North, model.East},
		{"Turn right from East", model.East, model.South},
		{"Turn right from South", model.South, model.West},
		{"Turn right from West", model.West, model.North},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := NewRover(0, 0, tt.initialDirection)
			initialPos := rover.Position

			rover.TurnRight()

			if rover.Direction != tt.expectedDirection {
				t.Errorf("Expected direction %s, got %s", tt.expectedDirection, rover.Direction)
			}

			// Position should remain unchanged
			if rover.Position != initialPos {
				t.Error("TurnRight should not change rover position")
			}
		})
	}
}

func TestGetPosition(t *testing.T) {
	rover := NewRover(10, 20, model.North)

	position := rover.GetPosition()

	if position.X != 10 || position.Y != 20 {
		t.Errorf("Expected position (10, 20), got (%d, %d)", position.X, position.Y)
	}
}

func TestGetDirection(t *testing.T) {
	rover := NewRover(0, 0, model.South)

	direction := rover.GetDirection()

	if direction != model.South {
		t.Errorf("Expected direction %s, got %s", model.South, direction)
	}
}

func TestComplexMovementSequence(t *testing.T) {
	rover := NewRover(0, 0, model.North)

	// Move forward, turn right, move forward, turn right, move forward, turn right, move forward
	// Should end up at origin facing West
	rover.Move()      // (0, 1) facing North
	rover.TurnRight() // (0, 1) facing East
	rover.Move()      // (1, 1) facing East
	rover.TurnRight() // (1, 1) facing South
	rover.Move()      // (1, 0) facing South
	rover.TurnRight() // (1, 0) facing West
	rover.Move()      // (0, 0) facing West

	if rover.Position.X != 0 || rover.Position.Y != 0 {
		t.Errorf("Expected final position (0, 0), got (%d, %d)", rover.Position.X, rover.Position.Y)
	}

	if rover.Direction != model.West {
		t.Errorf("Expected final direction %s, got %s", model.West, rover.Direction)
	}
}

func TestFullRotation(t *testing.T) {
	rover := NewRover(5, 5, model.North)
	initialPos := rover.Position

	// Four left turns should bring us back to North
	rover.TurnLeft() // West
	rover.TurnLeft() // South
	rover.TurnLeft() // East
	rover.TurnLeft() // North

	if rover.Direction != model.North {
		t.Errorf("After four left turns, expected direction %s, got %s", model.North, rover.Direction)
	}

	// Position should remain unchanged
	if rover.Position != initialPos {
		t.Error("Position should not change during rotation")
	}

	// Four right turns should also bring us back to North
	rover.TurnRight() // East
	rover.TurnRight() // South
	rover.TurnRight() // West
	rover.TurnRight() // North

	if rover.Direction != model.North {
		t.Errorf("After four right turns, expected direction %s, got %s", model.North, rover.Direction)
	}
}
