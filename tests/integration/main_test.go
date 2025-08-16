package integration

import (
	"os/exec"
	"strconv"
	"testing"
)

type testCase struct {
	name      string
	grid      int
	obstacles string
	commands  string
	want      string
}

func TestMarsRoverIntegration(t *testing.T) {
	tests := []testCase{
		{
			name:      "Success",
			grid:      5,
			obstacles: "[(1,2),(3,3)]",
			commands:  "LMLMLMLMM",
			want:      `{"final_position": [1, 3], "final_direction": "N", "status": "Success"}`,
		},
		// {
		// 	name:      "Obstacle encountered",
		// 	grid:      5,
		// 	obstacles: "[(1,2),(3,3)]",
		// 	commands:  "LMLMLMLMMMM",
		// 	want:      `{"final_position": [1, 2], "final_direction": "N", "status": "Obstacle encountered"}`,
		// },
		// {
		// 	name:      "Out of bounds",
		// 	grid:      5,
		// 	obstacles: "[]",
		// 	commands:  "MMMMMMMM",
		// 	want:      `{"final_position": [0, 4], "final_direction": "N", "status": "Out of bounds"}`,
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", "run", "../../src/main.go",
				"--grid_size", strconv.Itoa(tc.grid),
				"--obstacles", tc.obstacles,
				"--commands", tc.commands,
			)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("failed to run: %v, output: %s", err, out)
			}
			got := string(out)
			if got != tc.want {
				t.Errorf("got %s, want %s", got, tc.want)
			}
		})
	}
}

// [exmaple cli] `go run ./src/main.go --grid 5 --obstacles "[(1,2),(3,3)]" --commands "LMLMLMLMM"`

// [use-case]

// # =================
// # grid
// # size = 5
// # obstacles = [(1, 2), (3, 3)]
// # commands = "LMLMLMLMM"

// ----
// # Output:
// # {
// # "final_position": [1, 3],
// # "final_direction": "N",
// # "status": "Success"
// # }

// # =================

// # grid
// # size = 5
// # obstacles = [(1, 2), (3, 3)]
// # commands = "LMLMLMLMMMM"

// # --

// # {
// # "final_position": [1, 2],
// # "final_direction": "N",
// # "status": "Obstacle encountered"
// # }

// # =================

// # grid
// # size = 5
// # _
// # obstacles = []
// # commands = "MMMMMMMM"

// # --

// # {
// # "final_position": [0, 4],
// # "final_direction": "N",
// # "status": "Out of bounds"
// # }
