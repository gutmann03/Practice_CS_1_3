package maze

import (
	"math/rand"
	"time"
)

const (
	// Wall is used for all walls in the maze
	Wall bool = false

	// Empty is used for all space in the maze
	Empty bool = true
)

type MazeBuilder struct {
	level Level
}

func NewMazeBuilder(level Level) *MazeBuilder {
	return &MazeBuilder{
		level: level,
	}
}

func (b *MazeBuilder) GetLevel() Level {
	return b.level
}

func (b *MazeBuilder) CreateMaze() *Maze {
	width := START_SIZE + int(b.level)*STEP
	height := START_SIZE + int(b.level)*STEP
	maze := newMazePlan(width, height)
	return &Maze{
		Plan:         *maze,
		Width:        width,
		Height:       height,
		SelfLevel:    b.level,
		BlockSize:    BLOCK_SIZE - int(b.level),
		StartPoint:   getStartPoint(maze),
		EndPoint:     getEndPoint(maze),
		CurrentPoint: getStartPoint(maze),
	}
}

// New generates a Maze using Prim's Algorithm
func newMazePlan(w, h int) *MazePlan {
	maze := make([][]bool, w-2)
	rand.Seed(time.Now().Unix())

	for row := range maze {
		maze[row] = make([]bool, h)
		for ch := range maze[row] {
			maze[row][ch] = Wall
		}
	}

	p := &Point{X: rand.Intn(w - 2), Y: rand.Intn(h - 2)}
	maze[p.X][p.Y] = Empty

	var f *Point

	walls := adjacents(p, maze)

	for len(walls) > 0 {
		wall := walls[rand.Intn(len(walls))]

		for i, w := range walls {
			if w.X == wall.X && w.Y == wall.Y {
				walls = append(walls[:i], walls[i+1:]...)
				break
			}
		}

		opp := wall.opposite()

		if inMaze(opp.X, opp.Y, w-2, h-2) && !maze[opp.X][opp.Y] {
			maze[wall.X][wall.Y] = Empty
			maze[opp.X][opp.Y] = Empty
			walls = append(walls, adjacents(opp, maze)...)
			f = opp
		}
	}
	maze[f.X][f.Y] = Empty

	return borderedMaze((*MazePlan)(&maze))
}

func borderedMaze(mazeQ *MazePlan) *MazePlan {
	maze := *mazeQ
	b := make([][]bool, len(maze)+2)

	for r := range b {
		b[r] = make([]bool, len(maze[0]))

		for c := range b[r] {
			if r == 0 || r == len(maze)+1 || c == 0 || c == len(maze[0])+1 {
				b[r][c] = Wall
			} else {
				b[r][c] = maze[r-1][c-1]
			}
		}
	}

	return (*MazePlan)(&b)
}

func getStartPoint(mazeQ *MazePlan) (point Point) {
	maze := *mazeQ
	for h := range maze {
		for w := range maze[h] {
			if maze[h][len(maze[h])-w-1] {
				point = Point{X: len(maze[h]) - w - 1, Y: h}
				break
			}
		}
	}
	return
}

func getEndPoint(mazeQ *MazePlan) (point Point) {
	maze := *mazeQ
	for h := range maze {
		for w := range maze[len(maze[h])-h-1] {
			if maze[len(maze[h])-h-1][w] {
				point = Point{X: w, Y: len(maze[h]) - h - 1}
				break
			}
		}
	}
	return
}

func inMaze(x, y int, w, h int) bool {
	return x >= 0 && x < w && y >= 0 && y < h
}

// Width returns the width of the maze
func (m MazePlan) Width() int {
	return len(m)
}

// Height returns the height of the maze
func (m MazePlan) Height() int {
	return len(m[0])
}

func (m MazePlan) include(x, y int) bool {
	return x >= 0 && x < m.Width() && y >= 0 && y < m.Height()
}

func adjacents(p *Point, m MazePlan) []*Point {
	var res []*Point

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) || (i != 0 && j != 0) {
				continue
			}

			if !m.include(p.X+i, p.Y+j) {
				continue
			}

			if !m[p.X+i][p.Y+j] {
				res = append(res, &Point{p.X + i, p.Y + j, p})
			}
		}
	}

	return res
}
