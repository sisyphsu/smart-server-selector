package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os"
	"os/exec"
)

var sidebarWidth = 23
var exitFlag = 0
var app *tview.Application
var view *ServerUI

// Start the selector's render loop
func Start(a *tview.Application) {
	app = a
	view = newServersUI(loadServers())

	topFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(buildAboutUI(), sidebarWidth, 0, false).
		AddItem(buildSearchUI(), 0, 1, true)
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
		startVim() // open editor
	case tcell.KeyEnter:
		startSSH()
	}
	if exitFlag > 1 {
		app.Stop()
	}

	return event
}

// start vim subprocess
func startVim() {
	app.Suspend(func() {
		execute("vim", configFile)
		view.setServers(loadServers()) // reload
	})
}

// start ssh subprocess
func startSSH() {
	if view.offset >= len(view.visible) {
		return
	}
	app.Suspend(func() {
		s := view.visible[view.offset]
		execute("ssh", s.host)
	})
}

// execute the specified command
func execute(name string, args ...string) {
	// print command
	s := name
	if len(args) > 0 {
		for _, a := range args {
			s += " " + a
		}
	}
	println(os.Stdout, "> ", s)
	// start command
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	// print error
	if err != nil {
		println(os.Stdout, "exec error: ", err)
		println(os.Stdout, "press any key to continue")
		getchar()
	}
}
