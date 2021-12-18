package game

// Coordinates contains the coordinates of the snake's head
type Coordinates struct {
	X int
	Y int
}

// Speed reppresent the speed of the snake during the game
type Speed int8

//VERYSLOW tells the snake to go very slowly (500)
//SLOW tells the snake to go slowly (300ms)
//NORMAL tells the snake to go at normal speed (200ms)
//FAST tells the snake to go fast (100ms)
//VERYFAST tells the snake to go very fast (50ms)
const (
	VERYFAST Speed = 1
	FAST     Speed = 2
	NORMAL   Speed = 3
	SLOW     Speed = 4
	VERYSLOW Speed = 5
)

// Options reppresents the options of the game
type Options struct {
	PacmanEffect bool
	Speed        Speed
}

var termHeight = 0
var termWidth = 0

// Init initializes a new game
func Init(terminalWidth int, terminalHeight int) Snake {
	termWidth = terminalWidth
	termHeight = terminalHeight

	snake := Snake{
		Direction: RIGHT,
		Length:    0,
		Positions: []Coordinates{},
	}
	snake.Init()

	return snake
}
