package application

import (
	"container/heap"
	"fmt"
	"woolsocks-solution/internal/race-track/domain"
)

var directions = []domain.Node{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 0}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

type caseSolver struct{}

func NewCaseSolver() *caseSolver {
	return &caseSolver{}
}

// Solve method finds the least number of hops required to get from start to end point on the grid.
// It's an implementation of A* algorithm
func (a *caseSolver) Solve(grid [][]bool, start, end domain.Node, width, height int) string {
	initialState := &domain.State{Point: start, Velocity: domain.Node{}, Hops: 0, Heuristic: a.heuristic(start, end)}
	pq := &domain.PriorityQueue{initialState}
	heap.Init(pq)

	visited := make(map[domain.Node]map[domain.Node]bool)
	if visited[start] == nil {
		visited[start] = make(map[domain.Node]bool)
	}
	visited[start][domain.Node{}] = true

	// Perform a modified BFS using priority queue.
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*domain.State)

		if current.Point == end {
			return fmt.Sprintf("Optimal solution takes %d hops.", current.Hops)
		}

		// Iterate over all possible direction changes.
		for _, d := range directions {
			newV := domain.Node{X: current.Velocity.X + d.X, Y: current.Velocity.Y + d.Y}
			// Ensure the new velocity is within the allowed range.
			if newV.X < -3 || newV.X > 3 || newV.Y < -3 || newV.Y > 3 {
				continue
			}
			newP := domain.Node{X: current.Point.X + newV.X, Y: current.Point.Y + newV.Y}
			// Check if the new point is within the grid and not an obstacle.
			if a.isValid(newP.X, newP.Y, width, height, grid) {
				if visited[newP] == nil {
					visited[newP] = make(map[domain.Node]bool)
				}
				// If this state has not been visited yet, add it to the priority queue.
				if !visited[newP][newV] {
					newState := &domain.State{Point: newP, Velocity: newV, Hops: current.Hops + 1, Heuristic: a.heuristic(newP, end)}
					heap.Push(pq, newState)
					visited[newP][newV] = true
				}
			}
		}
	}

	return "No solution."
}

// heuristic calculates the Manhattan distance between two points.
func (a *caseSolver) heuristic(first, second domain.Node) int {
	return a.abs(first.X-second.X) + a.abs(first.Y-second.Y)
}

func (a *caseSolver) abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (a *caseSolver) isValid(x, y, width, height int, grid [][]bool) bool {
	return x >= 0 && x < width && y >= 0 && y < height && !grid[y][x]
}
