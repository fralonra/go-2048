package main

import (
	"log"
	"strconv"

	"github.com/fralonra/go-2048/core"
	ui "github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

const (
	size = 4
)

type app struct {
	game   *core.Game
	header *widgets.Paragraph
	table  *widgets.Table
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	a := &app{
		game: core.NewGame(),
	}

	a.header = widgets.NewParagraph()
	a.header.SetRect(0, 0, 20, 4)

	a.table = widgets.NewTable()
	a.table.Rows = [][]string{}
	a.table.TextStyle = ui.NewStyle(ui.ColorRed)
	a.table.TextAlignment = ui.AlignCenter
	a.table.SetRect(0, 6, 20, 15)

	a.renderNewGame()

	ui.Render(a.header)
	ui.Render(a.table)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "q", "<C-c>", "<Escape>":
				return
			case "r":
				a.game = core.NewGame()
				a.renderNewGame()
				ui.Render(a.header)
			case "<Left>":
				a.game.ToLeft()
			case "<Right>":
				a.game.ToRight()
			case "<Up>":
				a.game.ToTop()
			case "<Down>":
				a.game.ToBottom()
			}
			a.takeTurns()
		}
	}
}

func (a *app) takeTurns() {
	a.renderTableRows()
	switch a.game.State {
	case core.StateWin:
		{
			a.header.Text = "You Win!"
			ui.Render(a.header)
		}
	case core.StateLost:
		{
			a.header.Text = "You Lost! Max: " + strconv.Itoa(a.game.MaxNumber)
			ui.Render(a.header)
		}
	}
	ui.Render(a.table)
}

func (a *app) renderNewGame() {
	a.header.Text = "Play 2048 in cmd!"
	a.initTableRows()
	a.renderTableRows()
}

func (a *app) initTableRows() {
	for idx := 0; idx < core.Size; idx++ {
		row := a.game.GetRow(idx)
		displayRow := []string{}
		for _, item := range row {
			var text string
			if item > 0 {
				text = strconv.Itoa(item)
			} else {
				text = ""
			}
			displayRow = append(displayRow, text)
		}
		a.table.Rows = append(a.table.Rows, displayRow)
	}
}

func (a *app) renderTableRows() {
	for idx := 0; idx < core.Size; idx++ {
		row := a.game.GetRow(idx)
		for col, item := range row {
			var text string
			if item > 0 {
				text = strconv.Itoa(item)
			} else {
				text = ""
			}
			if idx == a.game.NewNumberPos[0] && col == a.game.NewNumberPos[1] {
				text += "*"
			}
			a.table.Rows[idx][col] = text
		}
	}
}
