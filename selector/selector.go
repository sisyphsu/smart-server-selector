package selector

import (
	"github.com/gdamore/tcell"
	ui "github.com/gizak/termui/v3"
	"github.com/rivo/tview"
	"os"
	"os/exec"
)

var app *tview.Application
var view *ServersUI

var exitFlag = 0
var serverActived *server

// Start the selector's render loop
func Start(a *tview.Application) {
	app = a
	view = newServersUI(loadServers())

	topFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(buildAboutUI(), sidebarWidth, 0, false).
		AddItem(buildKeywordUI(), 0, 1, true)
	btmFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(buildTipsUI(), sidebarWidth, 0, false).
		AddItem(view.flex, 0, 1, false)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topFlex, 3, 0, true).
		AddItem(btmFlex, 0, 1, false)

	app.SetInputCapture(onKeyEvent)
	app.SetRoot(flex, true)
}

func onKeyEvent(event *tcell.EventKey) *tcell.EventKey {
	if view.onEvent(event) {
		return nil
	}
	switch event.Key() {
	case tcell.KeyEscape:
		exitFlag += 1
	case tcell.KeyCtrlC:
		exitFlag += 9 // exit
	case tcell.KeyCtrlP:
		execute("vim", configFile) // open editor
	case tcell.KeyEnter:
		if serverActived != nil {
			execute("ssh", serverActived.host) // start ssh
		}
	default:
		// to keyword and server table
		//keyword.onEvent(e)
		//serverTable.onEvent(e)
		//ui.Render(buildAboutUI(), buildTipsUI(), keyword.build(), serverTable.build())
	}
	if exitFlag > 2 {
		app.Stop()
	}

	return event
}

// render global
func render() {
}

// execute the specified command
func execute(name string, args ...string) {
	app.Suspend(func() {
		cmd := exec.Command(name, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			exit(err)
		}
		if err = ui.Init(); err != nil {
			exit(err)
		}
	})
	// recover to Front
	render()
}

// exit
func exit(err error) {
	if exitFlag < 2 {
		return
	}
	if err != nil {
		println("error: ", err)
	}
	exitFlag = 2
}
