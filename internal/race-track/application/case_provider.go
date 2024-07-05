package application

import (
	"bufio"
	"fmt"
	"strings"
	"woolsocks-solution/internal/race-track/domain"
)

type caseProvider struct{}

func NewCaseProvider() *caseProvider {
	return &caseProvider{}
}

func (a *caseProvider) Get(input string) ([]domain.Case, error) {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)

	// Read the number of test cases
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read number of test cases")
	}

	var numTests int
	if _, err := fmt.Sscanf(scanner.Text(), "%d", &numTests); err != nil {
		return nil, fmt.Errorf("failed to parse number of test cases")
	}

	cases := make([]domain.Case, numTests)

	for i := 0; i < numTests; i++ {
		// Read grid width and height
		if !scanner.Scan() {
			return nil, fmt.Errorf("not enough information for test case %d: missing width and height", i+1)
		}
		var width, height int
		if _, err := fmt.Sscanf(scanner.Text(), "%d %d", &width, &height); err != nil {
			return nil, fmt.Errorf("failed to parse width and height for test case %d", i+1)
		}

		// Read start and end points
		if !scanner.Scan() {
			return nil, fmt.Errorf("not enough information for test case %d: missing start and end points", i+1)
		}
		var x1, y1, x2, y2 int
		if _, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &x1, &y1, &x2, &y2); err != nil {
			return nil, fmt.Errorf("failed to parse start and end points for test case %d", i+1)
		}
		start := domain.Node{X: x1, Y: y1}
		end := domain.Node{X: x2, Y: y2}

		// Read number of obstacles
		if !scanner.Scan() {
			return nil, fmt.Errorf("not enough information for test case %d: missing number of obstacles", i+1)
		}
		var numObstacles int
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &numObstacles); err != nil {
			return nil, fmt.Errorf("failed to parse number of obstacles for test case %d", i+1)
		}

		grid := make([][]bool, width)
		for j := range grid {
			grid[j] = make([]bool, height)
		}

		// Read each obstacle and mark the grid
		for j := 0; j < numObstacles; j++ {
			if !scanner.Scan() {
				return nil, fmt.Errorf("not enough information for obstacle %d in test case %d", j+1, i+1)
			}
			var x1, x2, y1, y2 int
			if _, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &x1, &x2, &y1, &y2); err != nil {
				return nil, fmt.Errorf("failed to parse obstacle %d for test case %d", j+1, i+1)
			}
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					if x >= width || y >= height {
						return nil, fmt.Errorf("obstacle %d for test case %d is out of grid bounds", j+1, i+1)
					}
					grid[x][y] = true
				}
			}
		}

		cases[i] = domain.Case{
			ID:           i + 1,
			Width:        width,
			Height:       height,
			Grid:         grid,
			Start:        start,
			End:          end,
			NumObstacles: numObstacles,
		}
	}

	// Check for extra lines in the input
	if scanner.Scan() {
		return nil, fmt.Errorf("extra data found after the declared number of test cases")
	}

	return cases, nil
}
