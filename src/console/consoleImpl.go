package console

import (
	"flag"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

type Modules struct {
	// TODO:
}

type consoleImpl struct {
	modules Modules
}

func Provide() *consoleImpl {
	// TODO: implement modules initialization
	return &consoleImpl{
		modules: Modules{},
	}
}

func (s *consoleImpl) Start() {
	log.Info("started")

	var gridSize int
	var obstaclesInput string
	var commands string

	flag.IntVar(&gridSize, "grid", 0, "Grid size")
	flag.StringVar(&obstaclesInput, "obstacles", "[]", "Obstacles in format [(x,y),(x,y),...]")
	flag.StringVar(&commands, "commands", "", "Commands string")
	flag.Parse()

	if gridSize == 0 {
		fmt.Println("Error: grid size is required")
		flag.Usage()
		return
	}

	if commands == "" {
		fmt.Println("Error: commands are required")
		flag.Usage()
		return
	}

	if err := s.validateObstaclesInput(obstaclesInput); err != nil {
		log.Error(err)
		return
	}

	obstacles := s.parseObstacles(obstaclesInput)

	fmt.Printf("Grid size: %d\n", gridSize)
	fmt.Printf("Obstacles: %v\n", obstacles)
	fmt.Printf("Commands: %s\n", commands)
}

func (s *consoleImpl) parseObstacles(obstaclesInput string) [][]int {
	var obstacles [][]int

	if obstaclesInput == "[]" || strings.TrimSpace(obstaclesInput[1:len(obstaclesInput)-1]) == "" {
		return obstacles
	}

	obstaclesInput = obstaclesInput[1 : len(obstaclesInput)-1] // remove outer [ ]
	pairs := obstaclesInput
	pairs = strings.ReplaceAll(pairs, "(", "")
	pairs = strings.ReplaceAll(pairs, ")", "")
	pairs = strings.ReplaceAll(pairs, " ", "")
	pairStrs := strings.Split(pairs, ",")

	for i := 0; i < len(pairStrs); i += 2 {
		if i+1 < len(pairStrs) {
			var x, y int
			fmt.Sscanf(pairStrs[i], "%d", &x)
			fmt.Sscanf(pairStrs[i+1], "%d", &y)
			obstacles = append(obstacles, []int{x, y})
		}
	}

	return obstacles
}

func (s *consoleImpl) validateObstaclesInput(obstaclesInput string) error {
	if len(obstaclesInput) < 2 {
		return fmt.Errorf("invalid obstacles format: too short")
	}

	if obstaclesInput[0] != '[' || obstaclesInput[len(obstaclesInput)-1] != ']' {
		return fmt.Errorf("invalid obstacles format: must be wrapped in []")
	}

	// Check if it's empty brackets
	content := obstaclesInput[1 : len(obstaclesInput)-1]
	if strings.TrimSpace(content) == "" {
		return nil // empty obstacles is valid
	}

	// Basic validation for parentheses pairs
	openCount := strings.Count(content, "(")
	closeCount := strings.Count(content, ")")
	if openCount != closeCount {
		return fmt.Errorf("invalid obstacles format: mismatched parentheses")
	}

	return nil
}
