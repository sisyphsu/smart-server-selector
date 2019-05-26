package core

import (
	"github.com/gizak/termui"
	"strings"
	"time"
)

// Keyword wrap keyword's input and render
type Keyword struct {
	text   string
	cursor string
	ticker *time.Ticker
	par    *termui.Par
}

func NewKeyword() *Keyword {
	k := new(Keyword)
	k.text = ""
	k.par = termui.NewPar("")
	k.par.X = sidebarWidth
	k.par.Height = 3
	k.par.Width = termui.TermWidth() - sidebarWidth
	k.par.TextFgColor = termui.ColorRed
	k.par.BorderLabel = "Search"
	k.par.BorderLabelFg = termui.ColorCyan
	k.par.BorderFg = termui.ColorCyan

	termui.Handle("/sys/kbd", k.onInput)

	k.flush()
	go k.setupCursorTimer()
	return k
}

// setup cursor's timer
func (k *Keyword) setupCursorTimer() {
	k.ticker = time.NewTicker(time.Second)
	time.Tick(time.Second)
	for range k.ticker.C {
		if len(k.cursor) == 0 {
			k.cursor = "▂"
		} else {
			k.cursor = ""
		}
		k.flush()
	}
}

// handle Keyword's input logic
func (k *Keyword) onInput(event termui.Event) {
	switch event.Path {
	case "/sys/kbd/<escape>":
		if l := len(k.text); l > 0 {
			k.setText("") // clear keyword
		}
	case "/sys/kbd/<backspace>", "/sys/kbd/<delete>", "/sys/kbd/C-8":
		if l := len(k.text); l > 0 {
			k.setText(k.text[:l-1]) // delete one keyword
		}
	default:
		parts := strings.Split(event.Path, "/")
		if l := len(parts); l > 0 && len(parts[l-1]) == 1 {
			k.cursor = ""
			k.setText(k.text + parts[l-1]) // input
		}
	}
}

// update keyword's text
func (k *Keyword) setText(text string) {
	k.text = text
	k.flush()
}

// flush redraw the keyword widget
func (k *Keyword) flush() {
	k.par.Text = k.text + k.cursor
	termui.Render(k.par)
}

func (k *Keyword) close() {
	k.ticker.Stop()
}
