package main

import termbox "github.com/nsf/termbox-go"

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
)

var (
	pointsChan         = make(chan int)
	keyboardEventsChan = make(chan keyboardEvent)
)

type game struct {
	board  *board
	score  int
	isOver bool
}

func initialScore() int {
	return 0
}

func initialBoard() *board {
	return newBoard(pointsChan, 19, 19)
}

func (g *game) end() {
	g.isOver = true
}

func (g *game) addPoints(p int) {
	g.score += p
}

// NewGame : returns a new game
func NewGame() *game {
	return &game{board: initialBoard(), score: initialScore()}
}

func (g *game) render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - g.board.width) / 2
		right  = (w + g.board.width) / 2
		top    = midY - (g.board.height / 2)
		bottom = midY + (g.board.height / 2) + 1
	)

	renderTitle(left, top)
	renderBoard(g.board, top, bottom, left)
	renderScore(left, bottom, g.score)
	renderQuitMessage(right, bottom)

	return termbox.Flush()
}

func (g *game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()

	go listenToKeyboard(keyboardEventsChan)

	if err := g.render(); err != nil {
		panic(err)
	}

mainloop:
	for {
		select {
		case p := <-pointsChan:
			g.addPoints(p)

		case e := <-keyboardEventsChan:
			switch e.eventType {
			case END:
				break mainloop
			}
		default:
		}
	}
}
