package selector

import (
	"github.com/gizak/termui"
)

// Handle all events
func HandleEvent(event termui.Event) (exit bool, update bool) {
	switch event.Path {
	case "/sys/kbd/<down>": // select down

	case "/sys/kbd/<up>": // select up

	case "/sys/kbd/<left>": // select ring left, if has left

	case "/sys/kbd/<right>": // select ring next, if has next

	case "/sys/kbd/<enter>": // confirm, ok

	case "/sys/kbd/<space>": // keyword + <space>

	case "/sys/kbd/<tab>": // select next option
	}
	// 2. start server list editor
	// 3. keyword input
	// 4. up/down select
	//println(event.Path)
	return
}
