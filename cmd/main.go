package main

import (
	"fmt"
	"sync"
	"woolsocks-solution/internal/race-track/application"
	"woolsocks-solution/internal/race-track/domain"
)

func main() {
	input := `2
5 5
4 0 4 4
1
1 4 2 3
3 3
0 0 2 2
2
1 1 0 2
0 2 1 1
`

	provider := application.NewCaseProvider()
	solver := application.NewCaseSolver()

	cases, err := provider.Get(input)
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup

	for _, testCase := range cases {
		wg.Add(1)
		go func(testCase domain.Case) {
			defer wg.Done()
			result := solver.Solve(testCase.Grid, testCase.Start, testCase.End, testCase.Width, testCase.Height)
			fmt.Printf("Case %d:\n%s\n", testCase.ID, result)
		}(testCase)
	}
	wg.Wait()
}
