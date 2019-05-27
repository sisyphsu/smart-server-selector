package selector

import (
	"github.com/gizak/termui"
	"strings"
	"time"
)

// Keyword wrap keyword's input and render
type Keyword struct {
	text     string
	cursor   string
	ticker   *time.Ticker
	par      *termui.Par
	onChange func(string)
}

func NewKeyword(onChange func(string)) *Keyword {
	k := new(Keyword)
	k.onChange = onChange
	k.text = ""
	k.par = termui.NewPar("")
	k.par.X = sidebarWidth
	k.par.Y = 0
	k.par.Width = termui.TermWidth() - sidebarWidth
	k.par.Height = 3
	k.par.TextFgColor = termui.ColorRed
	k.par.BorderLabel = "Search"
	k.par.BorderLabelFg = termui.ColorCyan
	k.par.BorderFg = termui.ColorCyan

	termui.Handle("/sys/kbd", k.onInput)

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
		k.render()
	}
}

// handle Keyword's input logic
func (k *Keyword) onInput(event termui.Event) {
	if !front {
		return
	}
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

// render keyword's text
func (k *Keyword) setText(text string) {
	k.text = text
	k.render()
	if k.onChange != nil {
		k.onChange(text)
	}
}

// render redraw the keyword widget
func (k *Keyword) render() {
	if front {
		k.par.Width = termui.TermWidth() - sidebarWidth
		k.par.Text = k.text + k.cursor
		termui.Render(k.par)
	}
}

func (k *Keyword) close() {
	k.ticker.Stop()
}
