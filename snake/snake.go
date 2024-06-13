package snake

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/rand"
)

const (
	size     = 30
	tickRate = 40 * time.Millisecond
)

type cell struct {
	x, y int
}

type direction uint8

const (
	up = direction(iota)
	down
	left
	right
)

type model struct {
	snake      []cell
	dir        direction
	dirChanged bool // dir was changed, but no affected snake position yet
	food       cell
	width      int
	height     int
}

func New(width, height int) model {
	m := model{
		snake: []cell{{x: size / 2, y: size / 2}},
		dir:   right,
	}
	m.food = spawnRandomFood(m)
	return m
}

type move struct{}

func (m model) Init() tea.Cmd {
	return tea.Tick(tickRate, func(t time.Time) tea.Msg {
		return move{}
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = (msg.Width - 2) / 2
		m.height = msg.Height - 2
	case move:
		snake := make([]cell, 1, len(m.snake))
		switch m.dir {
		case up:
			y := m.snake[0].y
			if y == 0 {
				y = m.height - 1
			} else {
				y--
			}
			snake[0] = cell{x: m.snake[0].x, y: y}
		case down:
			y := m.snake[0].y + 1
			if y >= m.height {
				y = 0
			}
			snake[0] = cell{x: m.snake[0].x, y: y}
		case left:
			x := m.snake[0].x
			if x == 0 {
				x = m.width - 1
			} else {
				x--
			}
			snake[0] = cell{x: x, y: m.snake[0].y}
		case right:
			x := m.snake[0].x + 1
			if x >= m.width {
				x = 0
			}
			snake[0] = cell{x: x, y: m.snake[0].y}
		}
		if cellIn(snake[0], m.snake) {
			return m, tea.Quit
		}
		if snake[0] == m.food {
			m.snake = append(snake, m.snake[0:len(m.snake)]...)
			m.food = spawnRandomFood(m)
		} else {
			m.snake = append(snake, m.snake[0:len(m.snake)-1]...)
		}
		m.dirChanged = false
		return m, tea.Tick(tickRate, func(t time.Time) tea.Msg {
			return move{}
		})
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		// Cycle between inputs
		case "up", "k", "w":
			if m.dir != down && !m.dirChanged {
				m.dirChanged = true
				m.dir = up
			}
			return m, nil
		// Cycle between inputs
		case "down", "j", "s":
			if m.dir != up && !m.dirChanged {
				m.dirChanged = true
				m.dir = down
			}
			return m, nil
		// Cycle between inputs
		case "left", "h", "a":
			if m.dir != right && !m.dirChanged {
				m.dirChanged = true
				m.dir = left
			}
			return m, nil
		// Cycle between inputs
		case "right", "l", "d":
			if m.dir != left && !m.dirChanged {
				m.dirChanged = true
				m.dir = right
			}
			return m, nil
		}
	}

	return m, cmd
}

func (m model) View() string {
	snakeBoard := []string{}

	for i := 0; i < m.height; i++ {
		row := []string{}
		for j := 0; j < m.width; j++ {
			c := cell{x: j, y: i}
			if c == m.food {
				row = append(
					row,
					lipgloss.
						NewStyle().
						Render(
							"🍎",
						),
				)
			} else if cellIn(c, m.snake) {
				row = append(
					row,
					lipgloss.
						NewStyle().
						Background(
							lipgloss.Color("#FFFFFF"),
						).
						Render("  "),
				)
			} else {
				row = append(row, "  ")
			}
		}
		snakeBoard = append(
			snakeBoard,
			lipgloss.
				JoinHorizontal(
					lipgloss.Top,
					row...,
				),
		)
	}

	return lipgloss.
		NewStyle().
		Border(lipgloss.RoundedBorder()).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				snakeBoard...,
			),
		)
}

// return true if c in arr
func cellIn(c cell, arr []cell) bool {
	for _, c2 := range arr {
		if c == c2 {
			return true
		}
	}
	return false
}

func getRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func (m model) getRandomCell() cell {
	return cell{x: getRandomInt(0, m.width), y: getRandomInt(0, m.height)}
}

func spawnRandomFood(m model) cell {
	if m.width == 0 || m.height == 0 {
		return cell{}
	}

	randCell := m.getRandomCell()
	if cellIn(randCell, m.snake) {
		return spawnRandomFood(m)
	}
	return randCell
}
