package selector

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"regexp"
	"strings"
)

var letter *regexp.Regexp

func init() {
	letter, _ = regexp.Compile("\\w")
}

// Keyword wrap keyword's input and render
type Keyword struct {
	text     string
	cursor   string
	input    *widgets.Paragraph
	onChange func(string)
}

func NewKeyword(onChange func(string)) *Keyword {
	k := new(Keyword)
	k.onChange = onChange
	k.text = ""
	k.input = widgets.NewParagraph()
	k.input.Title = "Search"
	k.input.TitleStyle.Fg = ui.ColorCyan
	k.input.Border = true
	k.input.Text = ""
	k.input.TextStyle.Fg = ui.ColorRed
	k.input.BorderStyle.Fg = ui.ColorCyan
	k.input.PaddingLeft = 1

	return k
}

// handle Keyword's input logic
func (k *Keyword) onEvent(event ui.Event) {
	if !Front {
		return
	}
	switch event.ID {
	case "<Escape>":
		if l := len(k.text); l > 0 {
			k.setText("") // clear keyword
		}
	case "<Backspace>", "<Delete>", "<C-8>":
		if l := len(k.text); l > 0 {
			k.setText(k.text[:l-1]) // delete one keyword
		}
	case "<Space>":
		k.cursor = ""
		k.setText(k.text + " ") // input
	default:
		parts := strings.Split(event.ID, "/")
		if l := len(parts); l > 0 && len(parts[l-1]) == 1 && letter.MatchString(parts[l-1]) {
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
	if Front {
		k.input.Text = k.text + k.cursor
		k.input.SetRect(sidebarWidth, 0, termWidth(), 3)
		ui.Render(k.input)
	}
}
