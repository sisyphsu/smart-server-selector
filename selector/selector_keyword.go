package selector

import (
	"github.com/gizak/termui/v3"
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
	input    *widgets.Paragraph
	onChange func(string)
}

func NewKeyword(onChange func(string)) *Keyword {
	k := new(Keyword)
	k.onChange = onChange
	k.text = ""
	k.input = widgets.NewParagraph()
	k.input.Title = "Search"
	k.input.TitleStyle.Fg = termui.ColorCyan
	k.input.Border = true
	k.input.Text = ""
	k.input.TextStyle.Fg = termui.ColorRed
	k.input.BorderStyle.Fg = termui.ColorCyan
	k.input.PaddingLeft = 1

	return k
}

// handle Keyword's input logic
func (k *Keyword) onEvent(event termui.Event) bool {
	if !Front {
		return false
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
		k.setText(k.text + " ") // input
	default:
		parts := strings.Split(event.ID, "/")
		if l := len(parts); l > 0 && len(parts[l-1]) == 1 && letter.MatchString(parts[l-1]) {
			k.setText(k.text + parts[l-1]) // input
		} else {
			return false
		}
	}
	return true
}

// render keyword's text
func (k *Keyword) setText(text string) {
	k.text = text
	if k.onChange != nil {
		k.onChange(text)
	}
}

func (k *Keyword) build() termui.Drawable {
	k.input.Text = k.text
	k.input.SetRect(sidebarWidth, 0, termWidth(), 3)
	return k.input
}
