package selector

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
	"os/exec"
)

const sidebarWidth = 23

var Exited bool
var Front bool
var keyword *Keyword
var serverTable *ServerTable

func Start() {
	serverTable = NewServerTable(loadServers())
	keyword = NewKeyword(serverTable.setKeyword)
	// init render
	render()

	// loop event
	uiEvents := ui.PollEvents()
	for !Exited {
		e := <-uiEvents
		if Front {
			switch e.ID {
			case "<C-c>":
				exit(nil)
			case "<Resize>":
				render()
			case "<C-p>":
				Front = false
				// start editor
				startEditor()
				// recover to Front
				keyword.setText("")
				render()
			default:
				keyword.onEvent(e)
				serverTable.onEvent(e)
			}
		}
	}

}

// render global
func render() {
	ui.Clear()

	Front = true
	renderAbout()
	renderHints()
	keyword.render()
	serverTable.render()
}

// render about of sidebar
func renderAbout() {
	about := widgets.NewParagraph()
	about.Title = "About"
	about.TitleStyle.Fg = ui.ColorCyan
	about.Text = "Smart Server Selector"
	about.TextStyle.Fg = ui.ColorCyan
	about.Border = true
	about.BorderStyle.Fg = ui.ColorCyan
	about.SetRect(0, 0, sidebarWidth, 3)

	ui.Render(about)
}

// render hints of sidebar
func renderHints() {
	hints := widgets.NewList()
	hints.Title = "Hints"
	hints.TitleStyle.Fg = ui.ColorCyan
	hints.Rows = hintsStr
	hints.TextStyle.Fg = ui.ColorYellow
	hints.Border = true
	hints.BorderStyle.Fg = ui.ColorCyan
	hints.SetRect(0, 3, sidebarWidth, termHeight())

	ui.Render(hints)
}

// start configuration's editor
func startEditor() {
	ui.Clear()

	cmd := exec.Command("vim", "~/test.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		exit(err)
	}
	if err = ui.Init(); err != nil {
		exit(err)
	}
}

// exit
func exit(err error) {
	if Exited {
		return
	}
	ui.Clear()
	if err != nil {
		println("error: ", err)
	}
	Exited = true
}
