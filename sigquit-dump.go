package sigquitdump

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT)

	go func(c <-chan os.Signal) {
		var buf [1048576]byte
		headLen := copy(buf[:], "\nSIGQUIT: begin of stack dump\n\n")
		for range c {
			n := headLen + runtime.Stack(buf[headLen:], true)
			n += copy(buf[n:], "\nSIGQUIT: end of stack dump\n")
			os.Stderr.Write(buf[:n])
		}
	}(c)
}
