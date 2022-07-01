package maze

const (
	START_SIZE int = 30
	STEP       int = 10
	BLOCK_SIZE int = 15
)

type Level int

const (
	Light Level = iota + 1
	SemiLight
	Medium
	SemiHard
	Hard
)
