package core

import (
	"math/rand"
	"time"
)

const (
	// The size of the game board.
	Size        = 4
	startNumber = 2
)

type GameState uint8

const (
	// The state represents the game is still being played.
	StateNormal GameState = iota
	// The state represents the player has won the game.
	StateWin
	// The state represents the player has lost.
	StateLost
)

var random *rand.Rand

type cell = int
type coordinate [2]int
type board [Size][Size]cell

// The struct stores the core data of the game.
type Game struct {
	// Max is the maximum number appeared in the board.
	MaxNumber int
	// NewNumberPos is the position where a new number is created this turn.
	NewNumberPos coordinate
	// State indicates the current state of the game.
	State GameState

	board
	availableCells []coordinate
}

// NewGame starts a new 2048 game. Returns the Game struct created.
func NewGame() *Game {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))

	game := &Game{}
	game.State = StateNormal
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			game.availableCells = append(game.availableCells, coordinate{row, col})
		}
	}
	game.MaxNumber = startNumber
	game.generateNewNumber()
	game.generateNewNumber()
	return game
}

// Slide to left.
func (g *Game) ToLeft() {
	hasMoved := g.mergeLeft()
	if hasMoved {
		g.State = g.newTurn()
	}
}

// Slide to right.
func (g *Game) ToRight() {
	hasMoved := g.mergeRight()
	if hasMoved {
		g.State = g.newTurn()
	}
}

// Slide up.
func (g *Game) ToTop() {
	hasMoved := g.mergeTop()
	if hasMoved {
		g.State = g.newTurn()
	}
}

// Slide down.
func (g *Game) ToBottom() {
	hasMoved := g.mergeBottom()
	if hasMoved {
		g.State = g.newTurn()
	}
}

// Get the tile number of the given position.
func (g *Game) Get(row int, colomn int) cell {
	return g.board[row][colomn]
}

// Get an array of tile numbers of the given row.
func (g *Game) GetRow(index int) [Size]cell {
	return g.board[index]
}

func (g *Game) mergeLeft() (hasMoved bool) {
	for idx, row := range g.board {
		moveState := false
		g.board[idx], moveState = merge(row)
		if moveState == true {
			hasMoved = moveState
		}
	}
	return
}

func (g *Game) mergeRight() (hasMoved bool) {
	for idx, row := range g.board {
		tmpRow, moveState := merge(reverseInts(row))
		g.board[idx] = reverseInts(tmpRow)
		if moveState == true {
			hasMoved = moveState
		}
	}
	return
}

func (g *Game) mergeTop() (hasMoved bool) {
	tmpRows := [Size][Size]cell{}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			tmpRows[row][col] = g.board[col][row]
		}
		moveState := false
		tmpRows[row], moveState = merge(tmpRows[row])
		if moveState == true {
			hasMoved = moveState
		}
	}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			g.board[row][col] = tmpRows[col][row]
		}
	}
	return
}

func (g *Game) mergeBottom() (hasMoved bool) {
	tmpRows := [Size][Size]cell{}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			tmpRows[row][col] = g.board[Size-1-col][row]
		}
		moveState := false
		tmpRows[row], moveState = merge(tmpRows[row])
		if moveState == true {
			hasMoved = moveState
		}
	}
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			g.board[row][col] = tmpRows[col][Size-1-row]
		}
	}
	return
}

func (g *Game) newTurn() GameState {
	if !g.canMove() {
		return StateLost
	}

	g.availableCells = []coordinate{}
	maxTmp := 0
	for i, row := range g.board {
		for j, item := range row {
			if item == 2048 {
				return StateWin
			}
			if item == 0 {
				g.availableCells = append(g.availableCells, coordinate{i, j})
			}
			if item > maxTmp {
				maxTmp = item
			}
		}
	}
	g.MaxNumber = maxTmp

	g.generateNewNumber()
	// TODO: no need to check all cells if they can move
	if !g.canMove() {
		return StateLost
	}

	return StateNormal
}

func (g *Game) generateNewNumber() {
	len := len(g.availableCells)
	if len > 0 {
		idx := random.Intn(len)
		g.NewNumberPos = g.availableCells[idx]
		g.board[g.NewNumberPos[0]][g.NewNumberPos[1]] = startNumber
		g.availableCells = append(g.availableCells[:idx], g.availableCells[idx+1:]...)
	}
}

func (g *Game) canMove() bool {
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			value := g.board[row][col]
			if value == 0 {
				return true
			}
			if col+1 < Size {
				if g.board[row][col+1] == value || g.board[row][col+1] == 0 {
					return true
				}
			}
			if row+1 < Size {
				if g.board[row+1][col] == value || g.board[row+1][col] == 0 {
					return true
				}
			}
		}
	}
	return false
}

func merge(row [Size]cell) ([Size]cell, bool) {
	currentPointer := 0
	hasMoved := false
	for idx, item := range row {
		if item == 0 || idx == 0 {
			continue
		}
		if item == row[currentPointer] {
			row[currentPointer] = item * 2
			row[idx] = 0
			currentPointer++
			hasMoved = true
		} else if row[currentPointer] == 0 {
			row[currentPointer] = item
			row[idx] = 0
			hasMoved = true
		} else {
			currentPointer++
			if idx != currentPointer {
				row[currentPointer] = item
				row[idx] = 0
				hasMoved = true
			}
		}
	}
	return row, hasMoved
}
