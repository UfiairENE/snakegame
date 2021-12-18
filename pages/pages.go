package pages

import (
	"fmt"
	"math"
	"github.com/ufiairene/snakegame/game"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

// Page is the type that reppresents the terminal page
type Page uint8

// WELCOME is the welcoming page
// TERMINALSMALL is the page that notifies the user that the terminal size is too small
// OPTIONS is the page that contains the game options
// GAME is the game page
const (
	WELCOME       Page = 0
	TERMINALSMALL Page = 1
	OPTIONS       Page = 2
	GAME          Page = 3
)

// TERMWIDTH is the minimum terminal width allowed to play
// TERMHEIGHT is the minimum terminal height allowed to play
const (
	TERMWIDTH  int = 80
	TERMHEIGHT int = 24
)

func drawBox() {
	termWidth, termHeight := termbox.Size()
	width := TERMWIDTH
	height := TERMHEIGHT
	if termWidth < width {
		width = termWidth
	}
	if termHeight < height {
		height = termHeight
	}
	pressESC := "Press ESC to exit"
	termbox.SetCell(0, 0, rune('┌'), termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(width-1, 0, rune('┐'), termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(0, height-1, rune('└'), termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(width-1, height-1, rune('┘'), termbox.ColorDefault, termbox.ColorDefault)
	for i := 1; i < width-1; i++ {
		if i-1 < len(pressESC) {
			termbox.SetCell(i, 0, rune(pressESC[i-1]), termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(i, 0, rune('─'), termbox.ColorDefault, termbox.ColorDefault)
		}
		termbox.SetCell(i, height-1, rune('─'), termbox.ColorDefault, termbox.ColorDefault)
	}
	for i := 1; i < height-1; i++ {
		termbox.SetCell(0, i, rune('│'), termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(width-1, i, rune('│'), termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
}

func drawString(x, y int, str string, fg, bg termbox.Attribute) {
	for pos, char := range str {
		termbox.SetCell(x+pos, y, rune(char), fg, bg)
	}
	termbox.Flush()
}

// DrawSmallTerminal shows a warning that tells the player that the terminal is too small
func DrawSmallTerminal(termWidth, termHeight int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBox()
	width := TERMWIDTH
	height := TERMHEIGHT
	if termWidth < width {
		width = termWidth
	}
	if termHeight < height {
		height = termHeight
	}
	str := "Terminal size too small"
	x := int(math.Floor(float64(width/2))) - int(len(str)/2)
	y := int(math.Floor(float64(height / 2)))
	drawString(x, y, str, termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

// DrawWelcome draws the initial page of the game
func DrawWelcome() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBox()
	termWidth, termHeight := termbox.Size()
	width := TERMWIDTH
	height := TERMHEIGHT
	if termWidth < width {
		width = termWidth
	}
	if termHeight < height {
		height = termHeight
	}
	var str string
	var x, y int
	str = "Snake"
	x = int(math.Floor(float64(width / 2)))
	y = int(math.Floor(float64(height / 2)))
	drawString(x-(len(str)/2), y-2, str, termbox.ColorDefault, termbox.ColorDefault)
	str = "Made by Mattia Costamagna - 2019"
	drawString(x-(len(str)/2), y, str, termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func getNextOption(currentOption string) string {
	switch currentOption {
	case "pacmaneffect":
		return fmt.Sprintf("speed_%d", game.VERYSLOW)
	case fmt.Sprintf("speed_%d", game.VERYSLOW):
		return fmt.Sprintf("speed_%d", game.SLOW)
	case fmt.Sprintf("speed_%d", game.SLOW):
		return fmt.Sprintf("speed_%d", game.NORMAL)
	case fmt.Sprintf("speed_%d", game.NORMAL):
		return fmt.Sprintf("speed_%d", game.FAST)
	case fmt.Sprintf("speed_%d", game.FAST):
		return fmt.Sprintf("speed_%d", game.VERYFAST)
	case fmt.Sprintf("speed_%d", game.VERYFAST):
		return "start"
	default:
		return "start"
	}
}

func getPrevOption(currentOption string) string {
	switch currentOption {
	case fmt.Sprintf("speed_%d", game.VERYSLOW):
		return "pacmaneffect"
	case fmt.Sprintf("speed_%d", game.SLOW):
		return fmt.Sprintf("speed_%d", game.VERYSLOW)
	case fmt.Sprintf("speed_%d", game.NORMAL):
		return fmt.Sprintf("speed_%d", game.SLOW)
	case fmt.Sprintf("speed_%d", game.FAST):
		return fmt.Sprintf("speed_%d", game.NORMAL)
	case fmt.Sprintf("speed_%d", game.VERYFAST):
		return fmt.Sprintf("speed_%d", game.FAST)
	case "start":
		return fmt.Sprintf("speed_%d", game.VERYFAST)
	default:
		return "pacmaneffect"
	}
}

// DrawOptions handles the options page shown before the game starts
func DrawOptions(keyPressed chan termbox.Key, gameOptions *game.Options, selectedOption string) bool {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBox()
	termWidth, termHeight := termbox.Size()
	width := TERMWIDTH
	height := TERMHEIGHT
	if termWidth < width {
		width = termWidth
	}
	if termHeight < height {
		height = termHeight
	}
	var str string
	var x, y int
	x = 30
	y = 5
	str = "Options"
	drawString(x, y, str, termbox.ColorDefault, termbox.ColorDefault)
	if gameOptions.PacmanEffect {
		str = "Pacman effect: [X]"
	} else {
		str = "Pacman effect: [ ]"
	}
	if selectedOption == "pacmaneffect" {
		drawString(x, y+2, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x, y+2, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	str = "Speed"
	if strings.HasPrefix(selectedOption, "speed") {
		drawString(x, y+3, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x, y+3, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	if gameOptions.Speed == game.VERYSLOW {
		str = "(*) Very slow"
	} else {
		str = "( ) Very slow"
	}
	if selectedOption == fmt.Sprintf("speed_%d", game.VERYSLOW) {
		drawString(x+1, y+4, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+4, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	if gameOptions.Speed == game.SLOW {
		str = "(*) Slow"
	} else {
		str = "( ) Slow"
	}
	if selectedOption == fmt.Sprintf("speed_%d", game.SLOW) {
		drawString(x+1, y+5, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+5, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	if gameOptions.Speed == game.NORMAL {
		str = "(*) Normal"
	} else {
		str = "( ) Normal"
	}
	if selectedOption == fmt.Sprintf("speed_%d", game.NORMAL) {
		drawString(x+1, y+6, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+6, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	if gameOptions.Speed == game.FAST {
		str = "(*) Fast"
	} else {
		str = "( ) Fast"
	}
	if selectedOption == fmt.Sprintf("speed_%d", game.FAST) {
		drawString(x+1, y+7, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+7, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	if gameOptions.Speed == game.VERYFAST {
		str = "(*) Very fast"
	} else {
		str = "( ) Very fast"
	}
	if selectedOption == fmt.Sprintf("speed_%d", game.VERYFAST) {
		drawString(x+1, y+8, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+8, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	str = "START"
	if selectedOption == "start" {
		drawString(x+1, y+10, str, termbox.ColorGreen, termbox.ColorDefault)
	} else {
		drawString(x+1, y+10, str, termbox.ColorDefault, termbox.ColorDefault)
	}
	termbox.Flush()
	for {
		key := <-keyPressed
		switch key {
		case termbox.KeyArrowDown:
			return DrawOptions(keyPressed, gameOptions, getNextOption(selectedOption))
		case termbox.KeyArrowUp:
			return DrawOptions(keyPressed, gameOptions, getPrevOption(selectedOption))
		case termbox.KeyEnter, termbox.KeySpace:
			switch selectedOption {
			case "start":
				return true
			case "pacmaneffect":
				gameOptions.PacmanEffect = !gameOptions.PacmanEffect
			case fmt.Sprintf("speed_%d", game.VERYSLOW):
				gameOptions.Speed = game.VERYSLOW
			case fmt.Sprintf("speed_%d", game.SLOW):
				gameOptions.Speed = game.SLOW
			case fmt.Sprintf("speed_%d", game.NORMAL):
				gameOptions.Speed = game.NORMAL
			case fmt.Sprintf("speed_%d", game.FAST):
				gameOptions.Speed = game.FAST
			case fmt.Sprintf("speed_%d", game.VERYFAST):
				gameOptions.Speed = game.VERYFAST
			}
			return DrawOptions(keyPressed, gameOptions, selectedOption)
		}
	}
}

func updateScore(score int) {
	str := fmt.Sprintf("Current score: %d", score)
	drawString(TERMWIDTH-1-len(str), TERMHEIGHT-1, str, termbox.ColorDefault, termbox.ColorDefault)
}

// Play handles the play page of the game
func Play(keyPressed chan termbox.Key, options game.Options) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	drawBox()
	updateScore(0)
	termWidth, termHeight := termbox.Size()
	width := TERMWIDTH
	height := TERMHEIGHT
	if termWidth < width {
		width = termWidth
	}
	if termHeight < height {
		height = termHeight
	}

	game.Init(width, height)
	snake := game.Init(width, height)

	food := game.Food{}
	food.Generate(snake)
	foodEaten := false
	alive := true
	paused := false
	maxLength := (width - 2) * (height - 2)
	for {
		if !alive {
			str := "You lost"
			x := int(math.Floor(float64(width / 2)))
			y := int(math.Floor(float64(height / 2)))
			drawString(x-(len(str)/2), y-1, str, termbox.ColorDefault, termbox.ColorDefault)
			str = "Press ENTER to start a new game"
			x = int(math.Floor(float64(width / 2)))
			y = int(math.Floor(float64(height/2)) + 2)
			drawString(x-(len(str)/2), y+1, str, termbox.ColorDefault, termbox.ColorDefault)
			break
		}
		select {
		case key := <-keyPressed:
			switch key {
			case termbox.KeyArrowUp:
				if snake.Direction != game.DOWN {
					snake.Direction = game.UP
				}
			case termbox.KeyArrowRight:
				if snake.Direction != game.LEFT {
					snake.Direction = game.RIGHT
				}
			case termbox.KeyArrowDown:
				if snake.Direction != game.UP {
					snake.Direction = game.DOWN
				}
			case termbox.KeyArrowLeft:
				if snake.Direction != game.RIGHT {
					snake.Direction = game.LEFT
				}
			case termbox.KeySpace:
				paused = !paused
			case termbox.KeyEnter:
				go func() { Play(keyPressed, options) }()
				return
			}
		default:
			if paused {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			if foodEaten {
				snake.Grow()
				alive = snake.Move(false, options)
				food.Generate(snake)

				updateScore(snake.Length - 1)
				if snake.Length == maxLength {
					str := "You win!"
					x := int(math.Floor(float64(width / 2)))
					y := int(math.Floor(float64(height / 2)))
					drawString(x-(len(str)/2), y-1, str, termbox.ColorDefault, termbox.ColorDefault)
					str = "Press ENTER to start a new game"
					x = int(math.Floor(float64(width / 2)))
					y = int(math.Floor(float64(height/2)) + 2)
					drawString(x-(len(str)/2), y+1, str, termbox.ColorDefault, termbox.ColorDefault)
					break
				}
			} else {
				alive = snake.Move(true, options)
			}
			foodEaten = false
			if snake.Positions[0].X == food.Position.X && snake.Positions[0].Y == food.Position.Y {
				foodEaten = true
			}
			time.Sleep(25 * time.Millisecond * time.Duration(options.Speed))
		}
	}
	for {
		key := <-keyPressed
		if key == termbox.KeyEnter {
			go func() { Play(keyPressed, options) }()
			return
		}
	}
}
