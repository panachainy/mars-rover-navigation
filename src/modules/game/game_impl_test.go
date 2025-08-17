package game

import (
	"mars-rover-navigation/src/model"
	"mars-rover-navigation/src/modules/environment"
	envMock "mars-rover-navigation/src/modules/environment/mock"
	"mars-rover-navigation/src/modules/rover"
	roverMock "mars-rover-navigation/src/modules/rover/mock"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game == nil {
		t.Error("NewGame() should return a non-nil game instance")
	}
	if game.envFactory == nil {
		t.Error("envFactory should not be nil")
	}
	if game.roverFactory == nil {
		t.Error("roverFactory should not be nil")
	}

	// Test that factories actually work and create proper instances
	testObstacles := []model.Position{{X: 1, Y: 1}}
	env := game.envFactory(5, testObstacles)
	if env == nil {
		t.Error("envFactory should create a non-nil environment")
	}

	rover := game.roverFactory(2, 3, model.Direction("E"))
	if rover == nil {
		t.Error("roverFactory should create a non-nil rover")
	}

	// Test that the rover is created with correct initial values
	if rover.GetPosition().X != 2 || rover.GetPosition().Y != 3 {
		t.Errorf("Expected rover position (2,3), got (%d,%d)", rover.GetPosition().X, rover.GetPosition().Y)
	}
	if rover.GetDirection() != model.Direction("E") {
		t.Errorf("Expected rover direction E, got %v", rover.GetDirection())
	}
}

func TestNewGame_FactoriesCreateDifferentInstances(t *testing.T) {
	game := NewGame()

	// Create two environments with same parameters
	env1 := game.envFactory(5, []model.Position{{X: 1, Y: 1}})
	env2 := game.envFactory(5, []model.Position{{X: 1, Y: 1}})

	// They should be different instances
	if env1 == env2 {
		t.Error("envFactory should create different instances on each call")
	}

	// Create two rovers with same parameters
	rover1 := game.roverFactory(0, 0, model.Direction("N"))
	rover2 := game.roverFactory(0, 0, model.Direction("N"))

	// They should be different instances
	if rover1 == rover2 {
		t.Error("roverFactory should create different instances on each call")
	}
}

func TestNewGameWithFactories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	envFactory := func(int, []model.Position) environment.Environment {
		return envMock.NewMockEnvironment(ctrl)
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return roverMock.NewMockRover(ctrl)
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	if game == nil {
		t.Error("NewGameWithFactories() should return a non-nil game instance")
	}
	if game.envFactory == nil {
		t.Error("envFactory should not be nil")
	}
	if game.roverFactory == nil {
		t.Error("roverFactory should not be nil")
	}
}

func TestIsValidInputs(t *testing.T) {
	tests := []struct {
		name      string
		size      int
		obstacles []model.Position
		commands  string
		expected  bool
	}{
		{
			name:      "valid inputs",
			size:      5,
			obstacles: []model.Position{{X: 1, Y: 1}, {X: 2, Y: 2}},
			commands:  "LRMM",
			expected:  true,
		},
		{
			name:      "zero size",
			size:      0,
			obstacles: []model.Position{},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "negative size",
			size:      -1,
			obstacles: []model.Position{},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "obstacle out of bounds - negative X",
			size:      5,
			obstacles: []model.Position{{X: -1, Y: 1}},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "obstacle out of bounds - negative Y",
			size:      5,
			obstacles: []model.Position{{X: 1, Y: -1}},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "obstacle out of bounds - X >= size",
			size:      5,
			obstacles: []model.Position{{X: 5, Y: 1}},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "obstacle out of bounds - Y >= size",
			size:      5,
			obstacles: []model.Position{{X: 1, Y: 5}},
			commands:  "LRM",
			expected:  false,
		},
		{
			name:      "invalid command character",
			size:      5,
			obstacles: []model.Position{},
			commands:  "LRMX",
			expected:  false,
		},
		{
			name:      "empty commands",
			size:      5,
			obstacles: []model.Position{},
			commands:  "",
			expected:  true,
		},
		{
			name:      "no obstacles",
			size:      5,
			obstacles: []model.Position{},
			commands:  "LRMM",
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidInputs(tt.size, tt.obstacles, tt.commands)
			if result != tt.expected {
				t.Errorf("isValidInputs() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestNavigateRover_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations
	mockRov.EXPECT().TurnLeft().Times(1)
	mockRov.EXPECT().TurnRight().Times(1)
	mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 0, Y: 1}).Times(2)
	mockEnv.EXPECT().CanMove(model.Position{X: 0, Y: 1}).Return(environment.Success).Times(2)
	mockRov.EXPECT().Move().Times(2)
	mockRov.EXPECT().GetPosition().Return(model.Position{X: 0, Y: 1})
	mockRov.EXPECT().GetDirection().Return(model.Direction("N"))

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(5, []model.Position{}, "LRMM")

	if result.Status != StatusSuccess {
		t.Errorf("Expected status %v, got %v", StatusSuccess, result.Status)
	}
}

func TestNavigateRover_ObstacleEncountered(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations
	mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 2, Y: 1})
	mockEnv.EXPECT().CanMove(model.Position{X: 2, Y: 1}).Return(environment.ObstacleEncountered)
	mockRov.EXPECT().GetPosition().Return(model.Position{X: 1, Y: 1})
	mockRov.EXPECT().GetDirection().Return(model.Direction("E"))

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(5, []model.Position{{X: 2, Y: 1}}, "M")

	if result.Status != StatusObstacleEncountered {
		t.Errorf("Expected status %v, got %v", StatusObstacleEncountered, result.Status)
	}
	expectedPosition := model.Position{X: 1, Y: 1}
	if result.FinalPosition != expectedPosition {
		t.Errorf("Expected position %v, got %v", expectedPosition, result.FinalPosition)
	}
	expectedDirection := model.Direction("E")
	if result.FinalDirection != expectedDirection {
		t.Errorf("Expected direction %v, got %v", expectedDirection, result.FinalDirection)
	}
}

func TestNavigateRover_OutOfBounds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations
	mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 4, Y: 5})
	mockEnv.EXPECT().CanMove(model.Position{X: 4, Y: 5}).Return(environment.OutOfBounds)
	mockRov.EXPECT().GetPosition().Return(model.Position{X: 4, Y: 4})
	mockRov.EXPECT().GetDirection().Return(model.Direction("N"))

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(5, []model.Position{}, "M")

	if result.Status != StatusOutOfBounds {
		t.Errorf("Expected status %v, got %v", StatusOutOfBounds, result.Status)
	}
	expectedPosition := model.Position{X: 4, Y: 4}
	if result.FinalPosition != expectedPosition {
		t.Errorf("Expected position %v, got %v", expectedPosition, result.FinalPosition)
	}
	expectedDirection := model.Direction("N")
	if result.FinalDirection != expectedDirection {
		t.Errorf("Expected direction %v, got %v", expectedDirection, result.FinalDirection)
	}
}

func TestNavigateRover_EmptyCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations - no commands should be called
	mockRov.EXPECT().GetPosition().Return(model.Position{X: 0, Y: 0})
	mockRov.EXPECT().GetDirection().Return(model.Direction("N"))

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(5, []model.Position{}, "")

	if result.Status != StatusSuccess {
		t.Errorf("Expected status %v, got %v", StatusSuccess, result.Status)
	}
}

func TestNavigateRover_ComplexScenario(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations for "LLRRMMM"
	mockRov.EXPECT().TurnLeft().Times(2)
	mockRov.EXPECT().TurnRight().Times(2)
	mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 2, Y: 1}).Times(3)
	mockEnv.EXPECT().CanMove(model.Position{X: 2, Y: 1}).Return(environment.Success).Times(3)
	mockRov.EXPECT().Move().Times(3)
	mockRov.EXPECT().GetPosition().Return(model.Position{X: 2, Y: 2})
	mockRov.EXPECT().GetDirection().Return(model.Direction("S"))

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(10, []model.Position{{X: 1, Y: 1}}, "LLRRMMM")

	if result.Status != StatusSuccess {
		t.Errorf("Expected status %v, got %v", StatusSuccess, result.Status)
	}
}

func TestNavigateRover_StopOnFirstObstacle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEnv := envMock.NewMockEnvironment(ctrl)
	mockRov := roverMock.NewMockRover(ctrl)

	// Set up expectations for "MMLR" - should stop after second M hits obstacle
	gomock.InOrder(
		// First M - success
		mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 1, Y: 2}),
		mockEnv.EXPECT().CanMove(model.Position{X: 1, Y: 2}).Return(environment.Success),
		mockRov.EXPECT().Move(),
		// Second M - obstacle encountered
		mockRov.EXPECT().GetTryMovePosition().Return(model.Position{X: 1, Y: 3}),
		mockEnv.EXPECT().CanMove(model.Position{X: 1, Y: 3}).Return(environment.ObstacleEncountered),
		// Return current position and direction when obstacle hit
		mockRov.EXPECT().GetPosition().Return(model.Position{X: 1, Y: 2}),
		mockRov.EXPECT().GetDirection().Return(model.Direction("N")),
	)

	envFactory := func(int, []model.Position) environment.Environment {
		return mockEnv
	}
	roverFactory := func(int, int, model.Direction) rover.Rover {
		return mockRov
	}

	game := NewGameWithFactories(envFactory, roverFactory)
	result := game.NavigateRover(5, []model.Position{{X: 1, Y: 3}}, "MMLR")

	if result.Status != StatusObstacleEncountered {
		t.Errorf("Expected status %v, got %v", StatusObstacleEncountered, result.Status)
	}
	expectedPosition := model.Position{X: 1, Y: 2}
	if result.FinalPosition != expectedPosition {
		t.Errorf("Expected position %v, got %v", expectedPosition, result.FinalPosition)
	}
}
