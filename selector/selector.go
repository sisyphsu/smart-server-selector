package selector

import (
	"fmt"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os"
	"os/exec"
)

var app *tview.Application
var view *ServersUI

var exitFlag = 0

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
		if view.offset < len(view.visible) {
			s := view.visible[view.offset]
			execute("ssh", s.host) // start ssh
		}
	}
	if exitFlag > 2 {
		app.Stop()
	}

	return event
}

// execute the specified command
func execute(name string, args ...string) {
	app.Suspend(func() {
		// print command
		s := name
		if len(args) > 0 {
			for _, a := range args {
				s += " " + a
			}
		}
		_, _ = fmt.Fprintln(os.Stdout, "> ", s)
		// start command
		cmd := exec.Command(name, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		// print error
		if err != nil {
			fmt.Fprintln(os.Stdout)
			fmt.Fprintln(os.Stdout, "exec error: ", err)
			fmt.Fprintln(os.Stdout, "press any key to continue")
			getchar()
		}
	})
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
