# SSHnake

<!--toc:start-->

- [SSHnake](#sshnake)
  - [Running the game](#running-the-game)
    - [Start the server](#start-the-server)
    - [Connect to the server](#connect-to-the-server)
  - [Gameplay](#gameplay)
  <!--toc:end-->

A simple single player snake game written in Go using [bubbletea](https://github.com/charmbracelet/bubbletea)
& Served over SSH using [Wish](https://github.com/charmbracelet/wish).

https://github.com/Gaurav-Gosain/sshnake/assets/55647468/799abd52-3c10-45c0-8ca6-716f3de28571

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
