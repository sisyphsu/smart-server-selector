package core

import (
	ui "github.com/gizak/termui/v3"
)

// Handle all events
func HandleEvent(event ui.Event) (exit bool, update bool) {
	switch event.ID {
	case "<down>":
		// select down
	case "<up>":
		// select up
	case "<left>":
		// select ring left, if has left
	case "<right>":
		// select ring next, if has next
	case "<enter>":
		// confirm, ok
	case "<space>":
		// keyword + <space>
	case "<tab>":
		// select next option
	case "<Escape>":
		// clear keyword
	case "<C-c>":
		exit = true
	case "<backspace>", "<delete>":
		// delete one keyword
	default:

	}
	// 2. start server list editor
	// 3. keyword input
	// 4. up/down select
	println(event.ID)
	return
}
