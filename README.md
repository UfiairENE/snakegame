The task is to implement a basic Golang version of the [Snake game] (https://en.wikipedia.org/wiki/Snake_(video_game_genre)) on the command line. Took about 4-5 hours to complete.


In order to play the game you'll need to build it first.<br/>
These are the steps you have to follow:

-   Install [go](https://golang.org/dl/) if you don't have it on your system already
-   Open a terminal
-   Go into the `snakegame` directory
-   Run this command

```bash
go build -ldflags="-s -w" -o snake main.go
```

-   This will create a `snake` binary that you can execute by running `./snake`

-   Use arrow keys to move the snake around
