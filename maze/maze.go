package maze

type MazeConfigurator interface {
	CreateMaze(level Level) *Maze
}

// Maze holds the bools of a generated maze
type MazePlan [][]bool

// Maze holds the square bool array of a generated maze and other parameters
type Maze struct {
	Plan       MazePlan
	Width      int
	Height     int
	SelfLevel  Level
	BlockSize  int
	StartPoint Point
	EndPoint   Point
}
