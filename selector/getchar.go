package selector

import (
	"os"
	"syscall"
	"unsafe"
)

// implement getChar
func getChar() {
	fd, err := os.Open("/dev/tty")
	if err != nil {
		panic(err)
	}

	var term syscall.Termios
	ptr := uintptr(unsafe.Pointer(&term))
	syscall.Syscall6(syscall.SYS_IOCTL, fd.Fd(), syscall.TIOCGETA, ptr, 0, 0, 0)

	// setup nobuffer
	term.Lflag &^= syscall.ICANON
	term.Cc[syscall.VMIN] = 1
	term.Cc[syscall.VTIME] = 0
	syscall.Syscall6(syscall.SYS_IOCTL, fd.Fd(), syscall.TIOCSETA, ptr, 0, 0, 0)

	// read
	b := make([]byte, 4)
	syscall.Read(int(fd.Fd()), b)
}
