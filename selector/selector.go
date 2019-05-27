package selector

import (
	ui "github.com/gizak/termui/v3"
	"os"
	"os/exec"
)

var Exited = false
var Front = true
var keyword *Keyword
var serverTable *ServerTable

func Start() {
	serverTable = NewServerTable(loadServers())
	keyword = NewKeyword(serverTable.setKeyword)
	// init render
	ui.Render(buildAbout(), buildHints(), keyword.build(), serverTable.build())

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
				startEditor()
			default:
				keyword.onEvent(e)
				serverTable.onEvent(e)
				ui.Render(buildAbout(), buildHints(), keyword.build(), serverTable.build())
			}
		}
	}

}

// render global
func render() {
	Front = true
	ui.Render(buildAbout(), buildHints(), keyword.build(), serverTable.build())
}

// start configuration's editor
func startEditor() {
	execute("vim", configFile)
}

// start ssh
func startSSH(s server) {
	execute("ssh", s.host)
}

// execute the specified command
func execute(name string, args ...string) {
	Front = false
	ui.Clear()
	ui.Render()
	println("> ", name, args)
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

	// recover to Front
	keyword.setText("")
	render()
}

// exit
func exit(err error) {
	if Exited {
		return
	}
	if err != nil {
		println("error: ", err)
	}
	Exited = true
}
