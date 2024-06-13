# SSHnake

A simple single player snake game written in Go using [bubbletea](https://github.com/charmbracelet/bubbletea)
& Served over SSH using [Wish](https://github.com/charmbracelet/wish).

> [!NOTE]
> Also supports vim-like movement keys (`hjkl`) and `ctrl+c` to quit the game.

## Running the game

### Start the server

```sh
go run main.go
```

### Connect to the server

```sh
ssh -p 23234 localhost
```

> [!NOTE]
> The server will listen on `localhost:23234`.

## Gameplay

The goal of the game is to guide your snake through the grid and eat the food
pellets without running into the walls or itself. Each time you eat a pellet,
your snake will grow in length.

You can use the following commands to control the snake:

- `w` (or up arrow): Move the snake up
- `a` (or left arrow): Move the snake left
- `s` (or down arrow): Move the snake down
- `d` (or right arrow): Move the snake right
- `q`: Quit the game
