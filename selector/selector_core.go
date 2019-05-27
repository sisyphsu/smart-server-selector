package selector

import (
	"github.com/gizak/termui"
	"github.com/nsf/termbox-go"
	"os"
	"os/exec"
)

const sidebarWidth = 23

var front bool
var keyword *Keyword
var serverTable *ServerTable

func Start() {
	serverTable = NewServerTable(loadServers())
	keyword = NewKeyword(serverTable.setKeyword)

	// register shortcut
	termui.Handle("/sys/wnd/resize", func(event termui.Event) {
		if front {
			render()
		}
	})
	termui.Handle("/sys/kbd/C-c", func(event termui.Event) {
		if front {
			termui.StopLoop() // exit
		}
	})
	termui.Handle("/sys/kbd/C-p", func(event termui.Event) {
		if front {
			front = false
			// start editor
			startEditor()
			// recover to front
			keyword.setText("")
			render()
		}
	})

	// init render
	render()
}

// render global
func render() {
	termui.Clear()

	front = true
	renderAbout()
	renderHints()
	keyword.render()
	serverTable.render()
}

// render about of sidebar
func renderAbout() {
	about := termui.NewPar("Smart Server Selector")
	about.X = 0
	about.Y = 0
	about.Width = sidebarWidth
	about.Height = 3
	about.TextFgColor = termui.ColorYellow
	about.BorderLabel = "About"
	about.BorderLabelFg = termui.ColorCyan
	about.BorderFg = termui.ColorCyan
	termui.Render(about)
}

// render hints of sidebar
func renderHints() {
	hints := termui.NewList()
	hints.Items = hintsStr
	hints.ItemFgColor = termui.ColorYellow
	hints.BorderLabel = "Hints"
	hints.BorderLabelFg = termui.ColorCyan
	hints.BorderFg = termui.ColorCyan
	hints.X = 0
	hints.Y = 3
	hints.Width = sidebarWidth
	hints.Height = termui.TermHeight() - 3
	termui.Render(hints)
}

// start configuration's editor
func startEditor() {
	termui.Clear()

	cmd := exec.Command("vim", "~/test.txt")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		exit(err)
	}
	if err = termbox.Init(); err != nil {
		exit(err)
	}
}

// exit
func exit(err error) {
	if err != nil {
		println("error: ", err)
	}
	termui.StopLoop()
	os.Exit(99)
}
