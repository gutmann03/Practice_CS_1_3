package maze

const (
	START_SIZE int = 30
	STEP       int = 4
	BLOCK_SIZE int = 20
)

type Level int

const (
	Easy Level = iota + 1
	Medium
	Hard
	Advanced
	Wizard
)
