package application

import (
	"testing"
	"woolsocks-solution/internal/race-track/domain"
)

func TestCaseSolver_Solve(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]bool
		start    domain.Node
		end      domain.Node
		width    int
		height   int
		expected string
	}{
		{
			name: "Sample input provided 1",
			grid: [][]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, true, true, true, true},
				{false, true, true, true, true},
				{false, false, false, false, false},
			},
			start:    domain.Node{X: 4, Y: 0},
			end:      domain.Node{X: 4, Y: 4},
			width:    5,
			height:   5,
			expected: "Optimal solution takes 7 hops.",
		},
		{
			name: "Sample input provided 2",
			grid: [][]bool{
				{false, true, false},
				{true, true, true},
				{false, true, false},
			},
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 2, Y: 2},
			width:    3,
			height:   3,
			expected: "No solution.",
		},
		{
			name: "Minimum Grid Size",
			grid: [][]bool{
				{false},
			},
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 0, Y: 0},
			width:    1,
			height:   1,
			expected: "Optimal solution takes 0 hops.",
		},
		{
			name: "Single Obstacle next to start",
			grid: [][]bool{
				{false, true},
				{false, false},
			},
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 1, Y: 1},
			width:    2,
			height:   2,
			expected: "Optimal solution takes 1 hops.",
		},
		{
			name: "Blocked Path",
			grid: [][]bool{
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 2, Y: 2},
			width:    3,
			height:   3,
			expected: "No solution.",
		},
		{
			name: "Path with Maximum Speed Constraints",
			grid: [][]bool{
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
				{false, false, false, false, false},
			},
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 4, Y: 4},
			width:    5,
			height:   5,
			expected: "Optimal solution takes 3 hops.",
		},
		{
			name: "Large Empty Grid",
			grid: func() [][]bool {
				grid := make([][]bool, 30)
				for i := range grid {
					grid[i] = make([]bool, 30)
				}
				return grid
			}(),
			start:    domain.Node{X: 0, Y: 0},
			end:      domain.Node{X: 29, Y: 29},
			width:    30,
			height:   30,
			expected: "Optimal solution takes 11 hops.",
		},
		{
			name: "Start Equals End",
			grid: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
			start:    domain.Node{X: 1, Y: 1},
			end:      domain.Node{X: 1, Y: 1},
			width:    3,
			height:   3,
			expected: "Optimal solution takes 0 hops.",
		},
	}

	solver := NewCaseSolver()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.grid != nil && len(test.grid) > 0 {
				for i := range test.grid {
					if test.grid[i] == nil {
						test.grid[i] = make([]bool, test.width)
					}
				}
			}
			result := solver.Solve(test.grid, test.start, test.end, test.width, test.height)
			if result != test.expected {
				t.Errorf("Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
