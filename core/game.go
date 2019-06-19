package core

import (
	"math/rand"
)

const (
	Size         = 4
	start_number = 2
)

type GameState uint8

const (
	StateNormal GameState = iota
	StateWin
	StateLost
)

type cell = int
type coordinate [2]int
type board [Size][Size]cell

type Game struct {
	board
	availableCells []coordinate
}

func NewGame() *Game {
	game := &Game{}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			game.availableCells = append(game.availableCells, coordinate{row, col})
		}
	}
	game.generateNewNumber()
	game.generateNewNumber()
	return game
}

func (g *Game) NewTurn() GameState {
	if len(g.availableCells) == 0 {
		return StateLost
	}
	g.availableCells = []coordinate{}
	for i, row := range g.board {
		for j, item := range row {
			if item == 2048 {
				return StateWin
			}
			if item == 0 {
				g.availableCells = append(g.availableCells, coordinate{i, j})
			}
		}
	}
	g.generateNewNumber()
	return StateNormal
}

func (g *Game) ToLeft() {
	for idx, row := range g.board {
		g.board[idx] = merge(row)
	}
}

func (g *Game) ToRight() {
	for idx, row := range g.board {
		g.board[idx] = reverseInts(merge(reverseInts(row)))
	}
}

func (g *Game) ToTop() {
	tmpRows := [Size][Size]cell{}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			tmpRows[row][col] = g.board[col][row]
		}
		tmpRows[row] = merge(tmpRows[row])
	}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			g.board[row][col] = tmpRows[col][row]
		}
	}
}

func (g *Game) ToBottom() {
	tmpRows := [Size][Size]cell{}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			tmpRows[row][col] = g.board[Size-1-col][row]
		}
		tmpRows[row] = merge(tmpRows[row])
	}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			g.board[row][col] = tmpRows[col][Size-1-row]
		}
	}
}

func (g *Game) Get(row int, colomn int) cell {
	return g.board[row][colomn]
}

func (g *Game) GetRow(index int) [Size]cell {
	return g.board[index]
}

func (g *Game) generateNewNumber() {
	idx := rand.Intn(len(g.availableCells))
	position := g.availableCells[idx]
	g.board[position[0]][position[1]] = start_number
	g.availableCells = append(g.availableCells[:idx], g.availableCells[idx+1:]...)
}

func merge(row [Size]cell) [Size]cell {
	currentPointer := 0
	for idx, item := range row {
		if item == 0 || idx == 0 {
			continue
		}
		if item == row[currentPointer] {
			row[currentPointer] = item * 2
			row[idx] = 0
			currentPointer++
		} else if row[currentPointer] == 0 {
			row[currentPointer] = item
			row[idx] = 0
		} else {
			currentPointer++
			if idx != currentPointer {
				row[currentPointer] = item
				row[idx] = 0
			}
		}
	}
	return row
}
