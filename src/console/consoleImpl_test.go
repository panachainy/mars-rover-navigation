package console

import (
	"bytes"
	"flag"
	"mars-rover-navigation/src/model"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestProvide(t *testing.T) {
	impl := Provide()
	if impl == nil {
		t.Error("Provide() returned nil")
	}
	if impl.modules == (Modules{}) {
		// Expected empty modules for now
	}
}

func TestConsoleImpl_Start_Success(t *testing.T) {
	// Save original args and restore after test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// Mock command line arguments
	os.Args = []string{"cmd", "-grid_size=5", "-obstacles=[(1,2),(3,3)]", "-commands=MMMRM"}
	
	// Reset flag package for testing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Capture stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	impl := Provide()
	impl.Start()

	// Restore stdout and read output
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verify output format
	if !strings.Contains(output, "final_position") {
		t.Errorf("Expected output to contain final_position, got: %s", output)
	}
}

func TestConsoleImpl_Start_MissingGridSize(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-commands=MMMRM"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Capture stderr
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	impl := Provide()
	impl.Start()

	w.Close()
	os.Stderr = oldStderr
	var buf bytes.Buffer
	buf.ReadFrom(r)
	// Test passes if no panic occurs
}

func TestConsoleImpl_Start_MissingCommands(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-grid_size=5"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	impl := Provide()
	impl.Start()
	// Test passes if no panic occurs
}

func TestConsoleImpl_ParseObstacles_Empty(t *testing.T) {
	impl := Provide()
	
	tests := []struct {
		name     string
		input    string
		expected []model.Position
	}{
		{
			name:     "empty brackets",
			input:    "[]",
			expected: []model.Position{},
		},
		{
			name:     "empty with spaces",
			input:    "[ ]",
			expected: []model.Position{},
		},
		{
			name:     "single obstacle",
			input:    "[(1,2)]",
			expected: []model.Position{{X: 1, Y: 2}},
		},
		{
			name:     "multiple obstacles",
			input:    "[(1,2),(3,4)]",
			expected: []model.Position{{X: 1, Y: 2}, {X: 3, Y: 4}},
		},
		{
			name:     "obstacles with spaces",
			input:    "[(1, 2), (3, 4)]",
			expected: []model.Position{{X: 1, Y: 2}, {X: 3, Y: 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := impl.parseObstacles(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("parseObstacles(%s) length = %d, want %d", tt.input, len(result), len(tt.expected))
				return
			}
			for i, pos := range result {
				if pos != tt.expected[i] {
					t.Errorf("parseObstacles(%s)[%d] = %v, want %v", tt.input, i, pos, tt.expected[i])
				}
			}
		})
	}
}

func TestConsoleImpl_ValidateObstaclesInput(t *testing.T) {
	impl := Provide()
	
	tests := []struct {
		name      string
		input     string
		expectErr bool
	}{
		{
			name:      "valid empty",
			input:     "[]",
			expectErr: false,
		},
		{
			name:      "valid with obstacles",
			input:     "[(1,2),(3,4)]",
			expectErr: false,
		},
		{
			name:      "too short",
			input:     "[",
			expectErr: true,
		},
		{
			name:      "missing opening bracket",
			input:     "1,2)]",
			expectErr: true,
		},
		{
			name:      "missing closing bracket",
			input:     "[(1,2)",
			expectErr: true,
		},
		{
			name:      "mismatched parentheses - extra open",
			input:     "[(1,2),((3,4)]",
			expectErr: true,
		},
		{
			name:      "mismatched parentheses - extra close",
			input:     "[(1,2),(3,4))]",
			expectErr: true,
		},
		{
			name:      "valid empty with spaces",
			input:     "[ ]",
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := impl.validateObstaclesInput(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("validateObstaclesInput(%s) error = %v, expectErr %v", tt.input, err, tt.expectErr)
			}
		})
	}
}

func TestConsoleImpl_ProcessFlags_Success(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-grid_size=5", "-obstacles=[(1,2)]", "-commands=MMMRM"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	impl := Provide()
	gridSize, obstacles, commands, err := impl.processFlags()

	if err != nil {
		t.Errorf("processFlags() error = %v, want nil", err)
	}
	if gridSize != 5 {
		t.Errorf("gridSize = %d, want 5", gridSize)
	}
	if commands != "MMMRM" {
		t.Errorf("commands = %s, want MMMRM", commands)
	}
	expectedObstacles := []model.Position{{X: 1, Y: 2}}
	if !reflect.DeepEqual(obstacles, expectedObstacles) {
		t.Errorf("obstacles = %v, want %v", obstacles, expectedObstacles)
	}
}

func TestConsoleImpl_ProcessFlags_InvalidObstacles(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "-grid_size=5", "-obstacles=invalid", "-commands=MMMRM"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	impl := Provide()
	_, _, _, err := impl.processFlags()

	if err == nil {
		t.Error("processFlags() error = nil, want error for invalid obstacles")
	}
}

func TestWire(t *testing.T) {
	console, err := Wire()
	if err != nil {
		t.Errorf("Wire() error = %v, want nil", err)
	}
	if console == nil {
		t.Error("Wire() returned nil console")
	}
	
	// Verify it implements Console interface
	_, ok := console.(Console)
	if !ok {
		t.Error("Wire() returned object that doesn't implement Console interface")
	}
}
