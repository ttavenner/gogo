package main

import (
	"fmt"

	runewidth "github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type board struct {
	height     int
	width      int
	pointsChan chan (int)
}

func newBoard(p chan (int), h, w int) *board {
	b := &board{
		height:     h,
		width:      w,
		pointsChan: p,
	}

	return b
}

// Cross character: U253C
// Dark circle: U25CF
// Light circle: U25CB
func renderBoard(b *board, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '|', defaultColor, bgColor)
		termbox.SetCell(left+b.width, i, '|', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+b.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+b.width, bottom, '┘', defaultColor, bgColor)

	fill(left, top, b.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, b.width, 1, termbox.Cell{Ch: '─'})

	fill(left, top+1, b.width, b.height-1, termbox.Cell{Ch: '┼'})
}

func renderScore(left, bottom, s int) {
	score := fmt.Sprintf("Score: %v", s)

	tbprint(left, bottom+1, defaultColor, defaultColor, score)
}

func renderQuitMessage(right, bottom int) {
	m := "Press ESC to quit"
	tbprint(right-17, bottom+1, defaultColor, defaultColor, m)
}

func renderTitle(left, top int) {
	tbprint(left, top-1, defaultColor, defaultColor, "Go")
}

func (b *board) addPoints(p int) {
	b.pointsChan <- p
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
