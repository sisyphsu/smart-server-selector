package selector

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"os"
	"os/exec"
)

var exited bool
var app *tview.Application
var view *ServerUI

// Start the selector's render loop
func Start(a *tview.Application) {
	app = a
	view = newServersUI(loadServers())

	topFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(buildAboutUI(), SidebarWidth, 0, false).
		AddItem(buildSearchUI(), 0, 1, true)
	btmFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(buildTipsUI(), SidebarWidth, 0, false).
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
		searchInput.SetText("")
	case tcell.KeyCtrlC, tcell.KeyCtrlD:
		exited = true // exit
	case tcell.KeyCtrlP:
		startVim() // open editor
	case tcell.KeyEnter:
		startSSH()
	}
	if exited {
		app.Stop()
	}

	return event
}

// start vim subprocess
func startVim() {
	app.Suspend(func() {
		execute("vim", SssFile)
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
		cmds := []string{SshOptions}
		if len(s.port) > 0 {
			cmds = append(cmds, "-p"+s.port)
		}
		if len(s.user) > 0 {
			cmds = append(cmds, s.user+"@"+s.host)
		} else {
			cmds = append(cmds, s.host)
		}
		execute("ssh", cmds...)
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
