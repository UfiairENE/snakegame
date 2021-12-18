package main

import (
	"github.com/ufiairene/snakegame/game"
	"github.com/ufiairene/snakegame/pages"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	changePage := make(chan pages.Page)
	keyPressed := make(chan termbox.Key)
	gameOptions := game.Options{
		PacmanEffect: true,
		Speed:        game.NORMAL,
	}
	go func() {
		for {
			currentPage := <-changePage
			switch currentPage {
			case pages.WELCOME:
				pages.DrawWelcome()
				go func() {
					time.Sleep(2 * time.Second)
					if currentPage != pages.TERMINALSMALL {
						width, height := termbox.Size()
						if width < pages.TERMWIDTH || height < pages.TERMHEIGHT {
							changePage <- pages.TERMINALSMALL
						} else {
							changePage <- pages.OPTIONS
						}
					}
				}()
			case pages.TERMINALSMALL:
				pages.DrawSmallTerminal(termbox.Size())
			case pages.OPTIONS:
				go func() {
					startPlay := pages.DrawOptions(keyPressed, &gameOptions, "start")
					if startPlay {
						changePage <- pages.GAME
					} else {
						changePage <- pages.OPTIONS
					}
				}()
			case pages.GAME:
				go pages.Play(keyPressed, gameOptions)
			}
		}
	}()
	changePage <- pages.WELCOME
	exit := false
	for {
		if exit {
			break
		}

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				exit = true
			case
				termbox.KeyArrowUp,
				termbox.KeyArrowDown,
				termbox.KeyArrowRight,
				termbox.KeyArrowLeft,
				termbox.KeyEnter,
				termbox.KeySpace:
				keyPressed <- ev.Key
			}

		case termbox.EventResize:
			width, height := termbox.Size()
			if width < pages.TERMWIDTH || height < pages.TERMHEIGHT {
				changePage <- pages.TERMINALSMALL
			} else {
				changePage <- pages.OPTIONS
			}
		}
	}
}
