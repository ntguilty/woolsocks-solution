package domain

type Case struct {
	ID           int
	Width        int
	Height       int
	Grid         [][]bool
	Start        Node
	End          Node
	NumObstacles int
}
