package main

import (
	"fmt"

	_ "github.com/m13253/go-sigquit-dump"
)

func main() {
	fmt.Println("Press Ctrl-\\ to print stack dump.")
	fmt.Println("Press Ctrl-C to terminate the program.")
	<-make(chan bool)
}
