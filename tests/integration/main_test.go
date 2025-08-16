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
		// FIXME: wait question answer
		// {
		// 	name:      "Success",
		// 	grid:      5,
		// 	obstacles: "[(1,2),(3,3)]",
		// 	commands:  "LMLMLMLMM",
		// 	want:      `{"final_position": [1, 3], "final_direction": "N", "status": "Success"}`,
		// },
		// {
		// 	name:      "Obstacle encountered",
		// 	grid:      5,
		// 	obstacles: "[(1,2),(3,3)]",
		// 	commands:  "LMLMLMLMMMM",
		// 	want:      `{"final_position": [1, 2], "final_direction": "N", "status": "Obstacle encountered"}`,
		// },
		{
			name:      "Out of bounds",
			grid:      5,
			obstacles: "[]",
			commands:  "MMMMMMMM",
			want:      "{\"final_position\": [0, 4], \"final_direction\": \"N\", \"status\": \"Out of bounds\"}\n",
		},
		{
			name:      "Invalid commands",
			grid:      5,
			obstacles: "[]",
			commands:  "LMXMLM",
			want:      "{\"final_position\": [0, 0], \"final_direction\": \"N\", \"status\": \"Invalid input\"}\n",
		},
		{
			name:      "Minimal grid 1x1",
			grid:      1,
			obstacles: "[]",
			commands:  "M",
			want:      "{\"final_position\": [0, 0], \"final_direction\": \"N\", \"status\": \"Out of bounds\"}\n",
		},
		{
			name:      "Minimal grid 1x1 turn only",
			grid:      1,
			obstacles: "[]",
			commands:  "LR",
			want:      "{\"final_position\": [0, 0], \"final_direction\": \"N\", \"status\": \"Success\"}\n",
		},
		{
			name:      "Large grid",
			grid:      100,
			obstacles: "[]",
			commands:  "RMMMMM",
			want:      "{\"final_position\": [5, 0], \"final_direction\": \"E\", \"status\": \"Success\"}\n",
		},
		{
			name:      "Long command string",
			grid:      10,
			obstacles: "[]",
			commands:  "RMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRMRM",
			want:      "{\"final_position\": [0, 0], \"final_direction\": \"N\", \"status\": \"Success\"}\n",
		},
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
