package application

import (
	"testing"
	"woolsocks-solution/internal/race-track/domain"
)

func TestCaseProvider_Get(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []domain.Case
		errMsg   string
	}{
		{
			name: "Single test case with obstacles",
			input: `1
5 5
4 0 4 4
1
1 4 2 3`,
			expected: []domain.Case{
				{
					ID:           1,
					Width:        5,
					Height:       5,
					Start:        domain.Node{X: 4, Y: 0},
					End:          domain.Node{X: 4, Y: 4},
					NumObstacles: 1,
					Grid: [][]bool{
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, true, true, true, true},
						{false, true, true, true, true},
						{false, false, false, false, false},
					},
				},
			},
			errMsg: "",
		},
		{
			name: "Multiple test cases",
			input: `2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1`,
			expected: []domain.Case{
				{
					ID:           1,
					Width:        5,
					Height:       5,
					Start:        domain.Node{X: 4, Y: 0},
					End:          domain.Node{X: 4, Y: 4},
					NumObstacles: 1,
					Grid: [][]bool{
						{false, false, false, false, false},
						{false, false, false, false, false},
						{false, true, true, true, true},
						{false, true, true, true, true},
						{false, false, false, false, false},
					},
				},
				{
					ID:           2,
					Width:        3,
					Height:       3,
					Start:        domain.Node{X: 0, Y: 0},
					End:          domain.Node{X: 2, Y: 2},
					NumObstacles: 2,
					Grid: [][]bool{
						{false, true, false},
						{true, true, true},
						{false, true, false},
					},
				},
			},
			errMsg: "",
		},
		{
			name: "Minimal grid size",
			input: `1
1 1
0 0 0 0
0`,
			expected: []domain.Case{
				{
					ID:           1,
					Width:        1,
					Height:       1,
					Start:        domain.Node{X: 0, Y: 0},
					End:          domain.Node{X: 0, Y: 0},
					NumObstacles: 0,
					Grid: [][]bool{
						{false},
					},
				},
			},
			errMsg: "",
		},
		{
			name: "Large empty grid",
			input: `1
30 30
0 0 29 29
0`,
			expected: []domain.Case{
				{
					ID:           1,
					Width:        30,
					Height:       30,
					Start:        domain.Node{X: 0, Y: 0},
					End:          domain.Node{X: 29, Y: 29},
					NumObstacles: 0,
					Grid: func() [][]bool {
						grid := make([][]bool, 30)
						for i := range grid {
							grid[i] = make([]bool, 30)
						}
						return grid
					}(),
				},
			},
			errMsg: "",
		},
		{
			name: "Blocked path",
			input: `1
3 3
0 0 2 2
1
0 2 1 1`,
			expected: []domain.Case{
				{
					ID:           1,
					Width:        3,
					Height:       3,
					Start:        domain.Node{X: 0, Y: 0},
					End:          domain.Node{X: 2, Y: 2},
					NumObstacles: 1,
					Grid: [][]bool{
						{false, false, false},
						{true, true, true},
						{false, false, false},
					},
				},
			},
			errMsg: "",
		},
		{
			name:   "No input",
			input:  ``,
			errMsg: "failed to read number of test cases",
		},
		{
			name: "Incomplete test case",
			input: `1
5 5`,
			errMsg: "not enough information for test case 1: missing start and end points",
		},
		{
			name: "Extra data after test cases",
			input: `1
5 5
4 0 4 4
0
extra data`,
			errMsg: "extra data found after the declared number of test cases",
		},
		{
			name: "Invalid obstacle coordinates",
			input: `1
3 3
0 0 2 2
1
0 4 1 1`,
			errMsg: "obstacle 1 for test case 1 is out of grid bounds",
		},
		{
			name: "Negative dimensions",
			input: `1
-5 5
0 0 4 4
0`,
			errMsg: "invalid grid dimensions for test case 1",
		},
		{
			name:   "Invalid number of test cases",
			input:  `a`,
			errMsg: "failed to parse number of test cases",
		},
		{
			name: "Non-integer coordinates",
			input: `1
5 5
a b c d
0`,
			errMsg: "failed to parse start and end points for test case 1",
		},
	}

	provider := NewCaseProvider()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cases, err := provider.Get(test.input)
			if err != nil {
				if err.Error() != test.errMsg {
					t.Errorf("expected error: %v, got: %v", test.errMsg, err.Error())
				}
				return
			}
			if test.errMsg != "" {
				t.Errorf("expected error: %v, got none", test.errMsg)
				return
			}
			if len(cases) != len(test.expected) {
				t.Fatalf("expected %d cases, got %d", len(test.expected), len(cases))
			}
			for i := range cases {
				if cases[i].ID != test.expected[i].ID ||
					cases[i].Width != test.expected[i].Width ||
					cases[i].Height != test.expected[i].Height ||
					cases[i].Start != test.expected[i].Start ||
					cases[i].End != test.expected[i].End ||
					cases[i].NumObstacles != test.expected[i].NumObstacles {
					t.Errorf("case %d: expected %+v, got %+v", i, test.expected[i], cases[i])
				}
				for y := range cases[i].Grid {
					for x := range cases[i].Grid[y] {
						if cases[i].Grid[y][x] != test.expected[i].Grid[y][x] {
							t.Errorf("case %d: grid mismatch at (%d,%d), expected %v, got %v",
								cases[i].ID, y, x, test.expected[i].Grid[y][x], cases[i].Grid[y][x])
						}
					}
				}
			}
		})
	}
}
